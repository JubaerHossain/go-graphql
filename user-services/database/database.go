package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	DB_HOST := "localhost"
	DB_PORT := "3306"
	DB_NAME := "go_crud"
	DB_PASS := "123"
	DB_USER := "root"
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &Database{db: db}
}

func (d *Database) Close() {
	if err := d.db.Close(); err != nil {
		log.Fatal(err)
	}
}
