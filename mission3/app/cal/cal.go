package cal

import (
	"database/sql"
	"fmt"
	"math"

	"com.mission/mission3/entity"
)

func FindInterestRate(db *sql.DB, dateInput entity.Date) (string, float64, entity.Date, entity.Date) {
	query := fmt.Sprintf(`SELECT * FROM promotion WHERE start_date <= '%v' AND end_date >= '%v';`, dateInput, dateInput)
	row := db.QueryRow(query)

	var promotion entity.Promotion
	err := row.Scan(&promotion.PromotionName, &promotion.Description, &promotion.StartDate, &promotion.EndDate)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	query = fmt.Sprintf(`SELECT * FROM rate WHERE promotion_name = '%s';`, promotion.PromotionName)
	row = db.QueryRow(query)

	var rate entity.Rate
	err = row.Scan(&rate.Rate, &rate.InterestRate, &rate.PromotionName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return promotion.PromotionName, rate.InterestRate, promotion.StartDate, promotion.EndDate
}

func CalculatePMT(input entity.PmtInput) float64 {
	PMT := input.DisbursementAmount / ((1 - (1 / math.Pow(1+(input.InterestRate/100/12), float64(input.NumberOfPayment)))) / (input.InterestRate / 100 / 12))
	PMT = math.Round(PMT*100) / 100
	return PMT
}
