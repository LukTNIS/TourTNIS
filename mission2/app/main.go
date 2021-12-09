package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5436
	username = "admin"
	password = "admin@1234"
	dbname   = "core_banking"
)

const accountCreateQuery = `
	CREATE TABLE account (
		account_number varchar(12) NOT NULL PRIMARY KEY,
		disbursement_amount FLOAT,
		number_of_payment INT,
		cal_date DATE
	);
`

const rateCreateQuery = `
	CREATE TABLE rate (
		rate varchar(20) NOT NULL PRIMARY KEY,
		interest_rate FLOAT,
		promotion_name varchar(20)
	);
`

const promotionCreateQuery = `
	CREATE TABLE promotion (
		promotion_name varchar(30) NOT NULL PRIMARY KEY,
		description varchar(50),
		start_date DATE,
		end_date DATE
	);
`

var db *sql.DB

func homePage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req request
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println("body: ", req.ReqBody)

	promotionName, interestRate := findInterestRate(req.ReqBody.CalDate)

	var pmtCal pmtCal
	pmtCal.DisbursementAmount = req.ReqBody.DisbursementAmount
	pmtCal.NumberOfPayment = req.ReqBody.NumberOfPayment
	pmtCal.InterestRate = interestRate

	var rs response
	var output []byte
	rs.RsBody.InstallmentAmount = calculatePMT(pmtCal)
	rs.RsBody.PromotionName = promotionName
	rs.RsBody.InterestRate = pmtCal.InterestRate
	rs.RsBody.AccountNumber = req.ReqBody.AccountNumber

	output, err = json.Marshal(rs)
	if err != nil {
		panic(err)
	}

	var accountNumber string
	query := fmt.Sprintf(`SELECT account_number FROM account WHERE account_number = '%s'`, rs.RsBody.AccountNumber)
	err = db.QueryRow(query).Scan(&accountNumber)
	if err != nil {
		query = fmt.Sprintf(`INSERT INTO account (account_number, disbursement_amount, number_of_payment, cal_date) VALUES ('%s', %v, %d, '%v');`, rs.RsBody.AccountNumber, req.ReqBody.DisbursementAmount, req.ReqBody.NumberOfPayment, req.ReqBody.CalDate)
		db.QueryRow(query)
	} else {
		query = fmt.Sprintf(`UPDATE account SET disbursement_amount = %v, number_of_payment = %d, cal_date = '%v' WHERE account_number = '%s'`, req.ReqBody.DisbursementAmount, req.ReqBody.NumberOfPayment, req.ReqBody.CalDate, rs.RsBody.AccountNumber)
		db.QueryRow(query)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(output))
}

func calculatePMT(input pmtCal) float64 {
	PMT := input.DisbursementAmount / ((1 - (1 / math.Pow(1+(input.InterestRate/100/12), float64(input.NumberOfPayment)))) / (input.InterestRate / 100 / 12))
	PMT = math.Round(PMT*100) / 100
	return PMT
}

func findInterestRate(dateInput Date) (string, float64) {
	query := fmt.Sprintf(`SELECT * FROM promotion WHERE start_date < '%v' AND end_date > '%v';`, dateInput, dateInput)
	row := db.QueryRow(query)

	var promotion promotion
	err := row.Scan(&promotion.PromotionName, &promotion.Description, &promotion.StartDate, &promotion.EndDate)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	query = fmt.Sprintf(`SELECT * FROM rate WHERE promotion_name = '%s';`, promotion.PromotionName)
	row = db.QueryRow(query)

	var rate rate
	err = row.Scan(&rate.Rate, &rate.InterestRate, &rate.PromotionName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return promotion.PromotionName, rate.InterestRate
}

func handleRequests() {
	http.HandleFunc("/dloan-payment/v1/calculate-installment-amount", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func connectDB(host string, port int, username string, password string, dbname string) *sql.DB {
	psqlInfo := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func createTableNecessary() {
	createTable(accountCreateQuery)
	createTable(rateCreateQuery)
	createTable(promotionCreateQuery)
}

func createTable(query string) {
	db.Exec(query)
}

func main() {
	db = connectDB(host, port, username, password, dbname)
	defer db.Close()
	createTableNecessary()

	handleRequests()
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

type Date time.Time

type request struct {
	ReqBody reqestBody `json:"req_body"`
}

type reqestBody struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	CalDate            Date    `json:"cal_date"`
	PaymentFrequency   int     `json:"payment_frequency"`
	PaymentUnit        string  `json:"payment_unit"`
	AccountNumber      string  `json:"account_number"`
}

type response struct {
	RsBody responseBody `json:"rs_body"`
}

type responseBody struct {
	InstallmentAmount float64 `json:"installment_amount"`
	PromotionName     string  `json:"promotion_name"`
	InterestRate      float64 `json:"interest_rate"`
	AccountNumber     string  `json:"account_number"`
}

type promotion struct {
	PromotionName string `json:"promotion_name"`
	Description   string `json:"description"`
	StartDate     Date   `json:"start_date"`
	EndDate       Date   `json:"end_date"`
}

type rate struct {
	Rate          string  `json:"rate"`
	InterestRate  float64 `json:"interest_rate"`
	PromotionName string  `json:"promotion_name"`
}

type pmtCal struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	InterestRate       float64 `json:"interest_rate"`
}
