package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	var req request
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	var rs response
	var output []byte
	rs.RsBody.InstallmentAmount = calculatePMT(req.ReqBody)
	output, err = json.Marshal(rs)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(output))
}

func calculatePMT(input reqestBody) float64 {
	PMT := input.DisbursementAmount / ((1 - (1 / math.Pow(1+(input.InterestRate/100/12), float64(input.NumberOfPayment)))) / (input.InterestRate / 100 / 12))
	PMT = math.Round(PMT*100) / 100
	return PMT
}

func handleRequests() {
	http.HandleFunc("/dloan-payment/v1/calculate-installment-amount", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}

type request struct {
	ReqBody reqestBody `json:"req_body"`
}

type reqestBody struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	InterestRate       float64 `json:"interest_rate"`
	PaymentFrequency   int     `json:"payment_frequency"`
	PaymentUnit        string  `json:"payment_unit"`
}

type response struct {
	RsBody responseBody `json:"rs_body"`
}

type responseBody struct {
	InstallmentAmount float64 `json:"installment_amount"`
}
