package query

import (
	"errors"
	"fmt"
	"lms/database"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func GetColumns(params graphql.ResolveParams) string {
	fieldASTs := params.Info.FieldASTs
	var fields = make(map[string]interface{})
	for _, val := range fieldASTs {
		var cols []string
		for _, sel := range val.SelectionSet.Selections {
			field, ok := sel.(*ast.Field)
			if ok {
				if field.Name.Kind == "Name" {
					cols = append(cols, field.Name.Value)
				}
			}
		}
		fields[val.Name.Value] = cols
	}

	funclabel := fmt.Sprint(params.Info.Path.Key)
	cols := fields[funclabel].([]string) //
	selectColumn := strings.Join(cols, ",")
	return selectColumn

}

func ModelColumn(selectColumn string, v interface{}) ([]interface{}, error) {
	var columns []interface{}
	for _, column := range strings.Split(selectColumn, ",") {
		fieldName := strings.ToTitle(column[:1]) + column[1:]
		fieldValue := reflect.ValueOf(v).Elem().FieldByName(fieldName)
		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("invalid field name: %s", fieldName)
		}
		columns = append(columns, fieldValue.Addr().Interface())
	}
	if len(columns) == 0 {
		return nil, errors.New("no columns selected")
	}
	return columns, nil
}

func QueryModel(modelType reflect.Type, modelName string, params graphql.ResolveParams) (interface{}, error) {
	// Get the database connection
	db := database.DB

	// Get the query parameters
	page, ok := params.Args["page"].(int)
	if !ok {
		page = 1
	}
	pageSize, ok := params.Args["pageSize"].(int)
	if !ok {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// Build the SQL query string
	selectColumn := GetColumns(params)

	sql := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d OFFSET %d;", selectColumn, modelName, pageSize, offset)

	// Execute the query
	rows, err := db.Query(sql)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	// Create a slice to hold the query results
	results := reflect.MakeSlice(reflect.SliceOf(modelType), 0, pageSize)

	// Loop through the query results and add them to the results slice
	for rows.Next() {
		// Create a new model instance
		model := reflect.New(modelType).Interface()

		// Get a list of pointers to the fields in the model struct
		columns, err := ModelColumn(selectColumn, model)
		if err != nil {
			return nil, err
		}

		// Scan the current row of data into the model struct fields
		err = rows.Scan(columns...)
		if err != nil {
			return nil, errors.New("no data found")
		}

		// Add the model to the results slice
		results = reflect.Append(results, reflect.ValueOf(model).Elem())
	}

	// Convert the results slice to an interface{} and return it
	return results.Interface(), nil
}
