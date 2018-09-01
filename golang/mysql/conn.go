package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func Init() {
	DB, err = sqlx.Connect("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/shoey?allowNativePasswords=true&parseTime=true")
	if err != nil {
		fmt.Println("show error: ", err)
	}
	DB.SetMaxOpenConns(0)
	DB.SetMaxIdleConns(1000)
	DB.SetConnMaxLifetime(0)
	err = DB.Ping()

}
