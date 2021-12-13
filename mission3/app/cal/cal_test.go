package cal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"testing"
	"time"

	"com.mission/mission3/entity"
	"github.com/DATA-DOG/go-sqlmock"
)

var promotion entity.Promotion

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCalculatePMT(t *testing.T) {
	var calDate entity.Date
	json.Unmarshal([]byte(`"2020-07-01"`), &calDate)

	type args struct {
		input entity.PmtInput
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "calculate installment_amount",
			args: args{
				input: entity.PmtInput{
					DisbursementAmount: 500000,
					NumberOfPayment:    48,
					InterestRate:       25,
					CalDate:            calDate,
				},
			},
			want: 16578.56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePMT(tt.args.input); got != tt.want {
				t.Errorf("CalculatePMT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindInterestRate(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	var calDateInput string = "2020-07-02"
	var calDate entity.Date
	json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, calDateInput)), &calDate)

	var startDate entity.Date
	json.Unmarshal([]byte(`"2020-07-01"`), &startDate)
	startDateMock, err := time.Parse("2006-01-02", "2020-07-01")
	if err != nil {
		panic(err)
	}

	var endDate entity.Date
	json.Unmarshal([]byte(`"2020-12-30"`), &endDate)
	endDateMock, err := time.Parse("2006-01-02", "2020-12-30")
	if err != nil {
		panic(err)
	}

	promotion = entity.Promotion{
		PromotionName: "Promo3",
		Description:   "Rate > 20 < 28",
		StartDate:     startDate,
		EndDate:       endDate,
	}
	rate := entity.Rate{
		Rate:          "RatePromo3",
		InterestRate:  25,
		PromotionName: "Promo3",
	}

	type args struct {
		db        *sql.DB
		dateInput entity.Date
	}
	tests := []struct {
		name    string
		args    args
		want    entity.FindInterestRateOutput
		wantErr bool
	}{
		{
			name: "Find Interest Rate is match.",
			args: args{db: db, dateInput: calDate},
			want: entity.FindInterestRateOutput{
				PromotionName: "Promo3",
				InterestRate:  25.0,
				StartDate:     startDate,
				EndDate:       endDate,
			},
			wantErr: false,
		},
		{
			name: "Find Interest Rate is not match.",
			args: args{db: db, dateInput: calDate},
			want: entity.FindInterestRateOutput{
				PromotionName: "Promo2",
				InterestRate:  25.0,
				StartDate:     startDate,
				EndDate:       endDate,
			},
			wantErr: true,
		},
		{
			name: "Find Interest Rate is error.",
			args: args{db: db, dateInput: calDate},
			want: entity.FindInterestRateOutput{
				PromotionName: "Promo3",
				InterestRate:  25.0,
				StartDate:     startDate,
			},
			wantErr: true,
		},
	}

	// fmt.Println(query, rows, query2, rows2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT * FROM promotion WHERE start_date <= '%s' AND end_date >= '%s';", calDateInput, calDateInput)

			rows := sqlmock.NewRows([]string{"promotion_name", "description", "start_date", "end_date"}).
				AddRow(promotion.PromotionName, promotion.Description, startDateMock, endDateMock)

			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			query2 := fmt.Sprintf("SELECT * FROM rate WHERE promotion_name = '%s';", tt.want.PromotionName)
			rows2 := sqlmock.NewRows([]string{"rate", "interest_rate", "promotion_name"}).AddRow(rate.Rate, rate.InterestRate, rate.PromotionName)
			mock.ExpectQuery(regexp.QuoteMeta(query2)).WillReturnRows(rows2)

			got, err := FindInterestRate(tt.args.db, tt.args.dateInput)
			fmt.Println(err != nil, tt.wantErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindInterestRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("FindInterestRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
