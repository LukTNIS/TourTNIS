package cal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

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

	var calDate entity.Date
	json.Unmarshal([]byte(`"2020-07-02"`), &calDate)

	var startDate entity.Date
	json.Unmarshal([]byte(`"2020-07-01"`), &startDate)

	var endDate entity.Date
	json.Unmarshal([]byte(`"2020-12-30"`), &endDate)

	promotion = entity.Promotion{
		PromotionName: "Promo3",
		Description:   "Rate > 20 < 28",
		StartDate:     startDate,
		EndDate:       endDate,
	}

	type args struct {
		db        *sql.DB
		dateInput entity.Date
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 float64
		want2 entity.Date
		want3 entity.Date
	}{
		{
			name:  "Find Interest Rate",
			args:  args{db: db, dateInput: calDate},
			want:  "Promo3",
			want1: 25.0,
			want2: startDate,
			want3: endDate,
		},
	}

	query := "SELECT * FROM promotion WHERE start_date <= '\\?' AND end_date >= '\\?';"

	fmt.Println("pass1")
	rows := sqlmock.NewRows([]string{"promotion_name", "description", "start_date", "end_date"}).
		AddRow(promotion.PromotionName, promotion.Description, "2020-07-01", "2020-12-30")

	fmt.Println("pass2")
	mock.ExpectQuery(query).WithArgs(promotion.PromotionName).WillReturnRows(rows)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := FindInterestRate(tt.args.db, tt.args.dateInput)
			fmt.Println("pass3")
			if got != tt.want {
				t.Errorf("FindInterestRate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindInterestRate() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("FindInterestRate() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("FindInterestRate() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
