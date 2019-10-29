package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

func updateData(conn *sql.DB, data interface{}) {

	sqlStr, sqlStrInsert := sqlString(data)

	res, err := conn.Exec(sqlStr)
	if err != nil {
		log.Println("update RainInfoHistory err:", err, sqlStr)

	}
	i, err := res.RowsAffected()
	if err != nil {
		log.Println("update RowsAffected RainInfoHistory err:", err, sqlStr)

	}
	if i == 0 {
		_, errInsert := conn.Exec(sqlStrInsert)
		if errInsert != nil {
			log.Println("insert RainInfoHistory err:", errInsert, sqlStrInsert)
		}
	}
}

func sqlString(data interface{}) (string, string) {
	var sqlStr string
	var sqlStrInsert string

	sqlStr = fmt.Sprintf(`update tabls_Name SET data1='%d', data2='%f', data3='%s'where keydata4='%s'`,
		0, 0.0, "", "")

	sqlStrInsert = fmt.Sprintf(`INSERT INTO tabls_Name (data1, data2, data3, keydata4) values('%d','%f','%s','%s') `,
		0, 0.0, "", "")

	return sqlStr, sqlStrInsert
}
