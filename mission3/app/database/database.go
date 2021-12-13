package database

import (
	"database/sql"
	"fmt"

	"com.mission/mission3/entity"

	_ "github.com/lib/pq"
)

func ConnectDB(dbConfig entity.ConfigDBInput) (*sql.DB, error) {
	sqlInfo := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname)

	fmt.Println("sqlInfo:", sqlInfo)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
