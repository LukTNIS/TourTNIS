package main2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"com.mission/mission3/cal"
	"com.mission/mission3/database"
	"com.mission/mission3/entity"
	"com.mission/mission3/formatinput"
	"com.mission/mission3/txt"
)

func RecordInstallmentAmountListByMonth(dbConfig entity.ConfigDBInput) error {
	dateInput := os.Args[1:][0]
	dateInputArr := strings.Split(dateInput, "-")
	dataDate := strings.Join(dateInputArr, "")

	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	startMonth := fmt.Sprintf(`%s-%s-01`, dateInputArr[0], dateInputArr[1])

	endMonthInt, err := strconv.Atoi(dateInputArr[1])
	if err != nil {
		return err
	}
	NextMonth := strconv.Itoa(endMonthInt + 1)
	if len(NextMonth) == 1 {
		NextMonth = "0" + NextMonth
	}
	endMonth := fmt.Sprintf(`%s-%s-01`, dateInputArr[0], NextMonth)
	query := fmt.Sprintf(`SELECT * FROM account WHERE cal_date >= '%s' AND cal_date <= '%s' ORDER BY cal_date ASC`, startMonth, endMonth)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var fileName string = formatinput.FileName(dataDate)
	os.Remove(fileName)
	err = txt.WriteHeader(fileName, dataDate)
	if err != nil {
		return err
	}

	var account entity.Account
	var totalRecord int
	var totalAmount float64
	for rows.Next() {
		err = rows.Scan(&account.AccountNumber, &account.DisbursementAmount, &account.NumberOfPayment, &account.CalDate)
		if err != nil {
			return err
		}
		findInterestRateOutput, err := cal.FindInterestRate(db, account.CalDate)
		if err != nil {
			return err
		}

		pmtInput := formatinput.PmtInput(account.DisbursementAmount, account.NumberOfPayment, findInterestRateOutput.InterestRate, account.CalDate)
		installmentAmount := cal.CalculatePMT(pmtInput)
		writeRecordInput := formatinput.WriteRecordInput(account.AccountNumber, installmentAmount, account.NumberOfPayment, findInterestRateOutput.InterestRate, findInterestRateOutput.PromotionName, findInterestRateOutput.StartDate, findInterestRateOutput.EndDate)

		err = txt.WriteRecord(fileName, writeRecordInput)
		if err != nil {
			return err
		}

		totalRecord++
		totalAmount += installmentAmount
	}

	err = txt.WriteFooter(fileName, totalRecord, totalAmount)
	if err != nil {
		return err
	}
	return nil
}
