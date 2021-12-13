package cal

import (
	"database/sql"
	"fmt"
	"math"

	"com.mission/mission3/entity"
)

func FindInterestRate(db *sql.DB, dateInput entity.Date) (entity.FindInterestRateOutput, error) {
	var output entity.FindInterestRateOutput

	fmt.Println("dateInput: ", dateInput)
	query := fmt.Sprintf(`SELECT * FROM promotion WHERE start_date <= '%v' AND end_date >= '%v';`, dateInput, dateInput)
	row := db.QueryRow(query)

	var promotion entity.Promotion
	err := row.Scan(&promotion.PromotionName, &promotion.Description, &promotion.StartDate, &promotion.EndDate)
	fmt.Println("=============", err)
	if err != nil && err != sql.ErrNoRows {
		return output, err
	}
	query = fmt.Sprintf(`SELECT * FROM rate WHERE promotion_name = '%s';`, promotion.PromotionName)
	row = db.QueryRow(query)

	var rate entity.Rate
	err = row.Scan(&rate.Rate, &rate.InterestRate, &rate.PromotionName)
	if err != nil && err != sql.ErrNoRows {
		return output, err
	}
	output = entity.FindInterestRateOutput{
		PromotionName: promotion.PromotionName,
		InterestRate:  rate.InterestRate,
		StartDate:     promotion.StartDate,
		EndDate:       promotion.EndDate,
	}

	return output, nil
}

func CalculatePMT(input entity.PmtInput) float64 {
	PMT := input.DisbursementAmount / ((1 - (1 / math.Pow(1+(input.InterestRate/100/12), float64(input.NumberOfPayment)))) / (input.InterestRate / 100 / 12))
	PMT = math.Round(PMT*100) / 100
	return PMT
}
