package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	configDB := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Addr:      "localhost:3306",
		DBName:    "task_db",
		Net:       "tcp",
		ParseTime: true,
	}

	db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}

}
