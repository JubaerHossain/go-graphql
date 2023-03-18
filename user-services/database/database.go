package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	DB_HOST := "localhost"
	DB_PORT := "3306"
	DB_NAME := "go_crud"
	DB_PASS := "123"
	DB_USER := "root"

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

func INSERT(tableName string, data interface{}) (int64, error) {
	var fields []string
	var values []string
	var args []interface{}

	// Get the type and value of the data parameter
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	// Determine if the data parameter is a struct or map
	var isMap bool
	if dataType.Kind() == reflect.Map {
		isMap = true
	} else if dataType.Kind() == reflect.Struct {
		isMap = false
	} else {
		return 0, fmt.Errorf("data parameter must be a struct or map[string]interface{}")
	}

	// Iterate over the fields of the data parameter and generate the SQL statement
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldName := field.Name

		var fieldValue reflect.Value
		if isMap {
			fieldValue = dataValue.MapIndex(reflect.ValueOf(fieldName))
		} else {
			fieldValue = dataValue.Field(i)
		}

		if fieldValue.IsValid() && fieldValue.Interface() != nil {
			fields = append(fields, fieldName)
			values = append(values, "?")
			args = append(args, fieldValue.Interface())
		}
	}

	// Generate the SQL statement
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(fields, ","), strings.Join(values, ","))

	// Execute the statement
	res, err := DB.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	// Get the ID of the inserted row
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
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
