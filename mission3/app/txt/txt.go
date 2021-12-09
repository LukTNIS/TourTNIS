package txt

import (
	"fmt"

	"com.mission/mission3/entity"
)

func WriteHeader(fileName, dataDate string) {
	textInput := fmt.Sprintf(`H|%s`, dataDate)
	WriteFileByLine(fileName, textInput)
}

func WriteRecord(fileName string, dataInput entity.WriteRecordInput) {
	textInput := fmt.Sprintf(`D|%s|%v|%d|%v|%s|%v|%v`, dataInput.AccountNumber, dataInput.InstallmentAmount, dataInput.NumberOfPayment, dataInput.InterestRate, dataInput.PromotionName, dataInput.StartDate, dataInput.EndDate)
	WriteFileByLine(fileName, textInput)
}

func WriteFooter(fileName string, totalRecord int, totalAmount float64) {
	textInput := fmt.Sprintf(`T|%d|%v`, totalRecord, totalAmount)
	WriteFileByLine(fileName, textInput)
}
