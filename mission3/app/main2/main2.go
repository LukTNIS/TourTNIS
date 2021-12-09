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

func RecordInstallmentAmountListByMonth(dbConfig entity.ConfigDBInput) {
	dateInput := os.Args[1:][0]
	dateInputArr := strings.Split(dateInput, "-")
	dataDate := strings.Join(dateInputArr, "")

	db := database.ConnectDB(dbConfig)
	defer db.Close()

	startMonth := fmt.Sprintf(`%s-%s-01`, dateInputArr[0], dateInputArr[1])

	endMonthInt, err := strconv.Atoi(dateInputArr[1])
	if err != nil {
		panic(err)
	}
	NextMonth := strconv.Itoa(endMonthInt + 1)
	if len(NextMonth) == 1 {
		NextMonth = "0" + NextMonth
	}
	endMonth := fmt.Sprintf(`%s-%s-01`, dateInputArr[0], NextMonth)
	query := fmt.Sprintf(`SELECT * FROM account WHERE cal_date >= '%s' AND cal_date <= '%s' ORDER BY cal_date ASC`, startMonth, endMonth)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("pass")
	defer rows.Close()

	var fileName string = formatinput.FileName(dataDate)
	txt.WriteHeader(fileName, dataDate)

	var account entity.Account
	var totalRecord int
	var totalAmount float64
	for rows.Next() {
		err = rows.Scan(&account.AccountNumber, &account.DisbursementAmount, &account.NumberOfPayment, &account.CalDate)
		if err != nil {
			panic(err)
		}
		promotionName, interestRate, startDate, endDate := cal.FindInterestRate(db, account.CalDate)

		pmtInput := formatinput.PmtInput(account.DisbursementAmount, account.NumberOfPayment, interestRate, account.CalDate)
		installmentAmount := cal.CalculatePMT(pmtInput)
		writeRecordInput := formatinput.WriteRecordInput(account.AccountNumber, installmentAmount, account.NumberOfPayment, interestRate, promotionName, startDate, endDate)

		txt.WriteRecord(fileName, writeRecordInput)

		totalRecord++
		totalAmount += installmentAmount
	}

	txt.WriteFooter(fileName, totalRecord, totalAmount)
}
