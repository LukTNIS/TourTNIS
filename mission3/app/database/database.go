package database

import (
	"database/sql"
	"fmt"

	"com.mission/mission3/entity"
)

func ConnectDB(dbConfig entity.ConfigDBInput) *sql.DB {
	psqlInfo := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname)

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
