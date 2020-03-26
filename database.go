package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/timetrack")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func insertAccount(account *Account) error {
	db := dbConnect()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO accounts (username, password) VALUES (?,?)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(account.Username, account.Password)
	return err
}

func lookupAccount(account *Account) error {
	db := dbConnect()
	defer db.Close()
	var id int
	stmt, err := db.Prepare("SELECT id FROM accounts WHERE username=? AND password=? LIMIT 1")

	if err != nil {
		panic(err)
	}
	fmt.Println("Query", account)
	res, err := stmt.Query(account.Username, account.Password)
	if res.Next() {
		err = res.Scan(&id)
	} else {
		fmt.Println("No results")
	}
	fmt.Println("Response", id)
	return err
}
