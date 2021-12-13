package main

import (
	"com.mission/mission3/formatinput"
	"com.mission/mission3/main2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5436
	username = "admin"
	password = "admin@1234"
	dbname   = "core_banking"
)

func main() {
	dbConfig := formatinput.ConfigDBInput(host, port, username, password, dbname)

	err := main2.RecordInstallmentAmountListByMonth(dbConfig)
	if err != nil {
		panic(err)
	}
}
