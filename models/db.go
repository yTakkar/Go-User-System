package models

import (
	"database/sql"

	// it is for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {
	db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
