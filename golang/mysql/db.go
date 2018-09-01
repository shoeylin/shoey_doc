package mysql

import (
	"fmt"
)

type db struct{}

func (p *db) dbtest(userId, userName, home, homeAddress string) error {
	tx, err := DB.Begin()
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				fmt.Printf("Could not roll back: %v\n", rollbackErr)
			}
		}
		if err := tx.Commit(); err != nil {
			fmt.Println("Commit error", err)
		}
	}()

	sql1 := "insert into user set id = ?, name = ?"
	sql2 := "insert into home set id = ?, address= ?"

	if err != nil {
		fmt.Println("dbtest error: ", err)
	}

	if _, err := tx.Exec(sql1, userId, userName); err != nil {
		fmt.Println("Exec error: ", err)
	}

	if _, err := tx.Exec(sql2, home, homeAddress); err != nil {
		fmt.Println("Exec error: ", err)
	}

	return err
}
