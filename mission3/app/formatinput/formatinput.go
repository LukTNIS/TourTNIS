package formatinput

import (
	"fmt"

	"com.mission/mission3/entity"
)

func FileName(dataDate string) string {
	fileName := fmt.Sprintf(`INSTALLMENT_%s.txt`, dataDate)
	return fileName
}

func PmtInput(disbursementAmount float64, numberOfPayment int, interestRate float64, calDate entity.Date) entity.PmtInput {
	pmtInput := entity.PmtInput{
		DisbursementAmount: disbursementAmount,
		NumberOfPayment:    numberOfPayment,
		InterestRate:       interestRate,
		CalDate:            calDate,
	}

	return pmtInput
}

func WriteRecordInput(accountNumber string, installmentAmount float64, numberOfPayment int, interestRate float64, promotionName string, startDate entity.Date, endDate entity.Date) entity.WriteRecordInput {
	writeRecordInput := entity.WriteRecordInput{
		AccountNumber:     accountNumber,
		InstallmentAmount: installmentAmount,
		NumberOfPayment:   numberOfPayment,
		InterestRate:      interestRate,
		PromotionName:     promotionName,
		StartDate:         startDate,
		EndDate:           endDate,
	}

	return writeRecordInput
}

func ConfigDBInput(host string, port int, username string, password string, dbname string) entity.ConfigDBInput {
	dbConfig := entity.ConfigDBInput{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Dbname:   dbname,
	}

	return dbConfig
}
