package query

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Select(tableName string, data map[string]interface{},  where string, id int, db *sql.DB) (*sql.Rows, error) {
	// Generate the SQL statement
	fmt.Println(data)
	sql := "SELECT * FROM " + tableName + " WHERE " + where

	// Use prepared statement to execute the statement
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func RowExistsByID(table string, id int64, db *sql.DB) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE id = ?)", table)
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func Update(form map[string]interface{}, db *sql.DB) error {
	id, _ := form["id"].(int)
	if !RowExistsByID(form["table"].(string), int64(id), db) {
		return errors.New("id not found")
	}

	table, _ := form["table"].(string)
	if table == "" {
		return errors.New("table is missing")
	}

	// value update
	var valueUpdate []string
	for key, val := range form {
		if key != "id" && key != "table" {
			valueUpdate = append(valueUpdate, fmt.Sprintf("%s = '%s'", key, val))
		}
	}

	// query
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", table, strings.Join(valueUpdate, ", "), id)

	// execute query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func Insert(tableName string, data interface{}, db *sql.DB) (int64, error) {
	// Get the type and value of the data parameter
	dataType, dataValue := reflect.TypeOf(data), reflect.ValueOf(data)

	// Determine if the data parameter is a struct or map
	var isMap bool
	if dataType.Kind() == reflect.Map {
		isMap = true
	} else if dataType.Kind() == reflect.Struct {
		isMap = false
	} else {
		return 0, fmt.Errorf("data parameter must be a struct or map[string]interface{}")
	}

	// Cache the type information
	var buf bytes.Buffer
	if isMap {
		dataType = reflect.TypeOf(map[string]interface{}{})
		buf.WriteString("INSERT INTO " + tableName + " (")
	} else {
		buf.WriteString("INSERT INTO " + tableName + " (")
		dataType = reflect.TypeOf(data)
	}

	// Iterate over the fields of the data parameter and generate the SQL statement
	var fields []string
	var values []string
	var args []interface{}
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

	buf.WriteString(strings.Join(fields, ","))
	buf.WriteString(") VALUES (")
	buf.WriteString(strings.Join(values, ","))
	buf.WriteString(")")

	// Generate the SQL statement
	sql := buf.String()

	// Use prepared statement to execute the statement
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
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
