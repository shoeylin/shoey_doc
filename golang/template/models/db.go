package models

import (
	"database/sql"
	"log"
)

var ConnAry []*sql.DB

func GetDBConn(i int, connStr string) {

	var err error
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error(), " db: ", i)
		return
	}
	errPing := db.Ping()
	if errPing != nil {
		log.Println("db ", i, " errPing :", errPing)
		return
	}

	ConnAry = append(ConnAry, db)
	log.Println("DB Connect OK ", i)
}
