package query

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func Select(tableName string, data map[string]interface{}, where string, id int, db *sql.DB) (*sql.Rows, error) {
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

func Insert(form map[string]interface{}, db *sql.DB) (int64, error) {

	table := form["table"].(string)
	if table == "" {
		return 0, errors.New("table is missing")
	}

	// value insert
	var valueInsert []string
	for key, val := range form {
		if key != "id" && key != "table" {
			valueInsert = append(valueInsert, fmt.Sprintf("'%s'", val))
		}
	}

	// query

	query := fmt.Sprintf("INSERT INTO %s VALUES (NULL, %s)", table, strings.Join(valueInsert, ", "))
	// execute query

	_, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	// get last insert id
	var lastInsertID int64
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil

}
