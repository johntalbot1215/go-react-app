package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/timetrack")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertAccount(account *Account) error {
	db := dbConnect()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO accounts (username, password) VALUES (?,?)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(account.Username, account.Password)
	return err
}
