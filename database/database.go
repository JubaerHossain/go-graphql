package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"lms/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	DB_HOST := config.Env("DB_HOST")
	DB_PORT := config.Env("DB_PORT")
	DB_NAME := config.Env("DB_DATABASE")
	DB_PASS := config.Env("DB_PASSWORD")
	DB_USER := config.Env("DB_USERNAME")

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	fmt.Println("Connected to database!")
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
