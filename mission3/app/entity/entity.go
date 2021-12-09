package entity

import (
	"encoding/json"
	"time"
)

type Date time.Time

type Account struct {
	AccountNumber      string  `json:"account_number"`
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	CalDate            Date    `json:"cal_date"`
}

type Promotion struct {
	PromotionName string `json:"promotion_name"`
	Description   string `json:"description"`
	StartDate     Date   `json:"start_date"`
	EndDate       Date   `json:"end_date"`
}

type Rate struct {
	Rate          string  `json:"rate"`
	InterestRate  float64 `json:"interest_rate"`
	PromotionName string  `json:"promotion_name"`
}

type PmtInput struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	InterestRate       float64 `json:"interest_rate"`
	CalDate            Date    `json:"cal_date"`
}

type WriteRecordInput struct {
	AccountNumber     string  `json:"account_number"`
	InstallmentAmount float64 `json:"installment_amount"`
	NumberOfPayment   int     `json:"number_of_payment"`
	InterestRate      float64 `json:"interest_rate"`
	PromotionName     string  `json:"promotion_name"`
	StartDate         Date    `json:"start_date"`
	EndDate           Date    `json:"end_date"`
}

type ConfigDBInput struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format(`"2006-01-02"`)), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	tm, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return json.Unmarshal(b, (*time.Time)(d))
	}
	*d = Date(tm)
	return nil
}
