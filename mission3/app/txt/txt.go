package txt

import (
	"fmt"

	"com.mission/mission3/entity"
)

func WriteHeader(fileName, dataDate string) error {
	textInput := fmt.Sprintf(`H|%s`, dataDate)
	err := WriteFileByLine(fileName, textInput)
	if err != nil {
		return err
	}
	return nil
}

func WriteRecord(fileName string, dataInput entity.WriteRecordInput) error {
	textInput := fmt.Sprintf(`D|%s|%v|%d|%v|%s|%v|%v`, dataInput.AccountNumber, dataInput.InstallmentAmount, dataInput.NumberOfPayment, dataInput.InterestRate, dataInput.PromotionName, dataInput.StartDate, dataInput.EndDate)
	err := WriteFileByLine(fileName, textInput)
	if err != nil {
		return err
	}
	return nil
}

func WriteFooter(fileName string, totalRecord int, totalAmount float64) error {
	textInput := fmt.Sprintf(`T|%d|%v`, totalRecord, totalAmount)
	err := WriteFileByLine(fileName, textInput)
	if err != nil {
		return err
	}
	return nil
}
