package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/golang_db?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)  // minimal connection pool
	db.SetMaxOpenConns(100) // max connection pool
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
