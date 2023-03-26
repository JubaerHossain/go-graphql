package resolver

import (
	"errors"
	"fmt"
	"lms/database"
	"lms/graphql/types"
	"lms/query"
	"lms/utils"
	"strings"

	"github.com/graphql-go/graphql"
)

func GetUsers(params graphql.ResolveParams) (interface{}, error) {
	page, ok := params.Args["page"].(int)
	if !ok {
		page = 1
	}
	pageSize, ok := params.Args["pageSize"].(int)
	if !ok {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	colmap := query.TableFields(params.Info.FieldASTs)
	funclabel := fmt.Sprint(params.Info.Path.Key)
	cols := colmap[funclabel].([]string) //
	selectColumn := strings.Join(cols, ",")
	sql := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d OFFSET %d;", selectColumn, "users", pageSize, offset)

	fmt.Println(sql)
	rows, err := query.GetAllRowsByQuery(sql, database.DB)
	if err != nil {
		return nil, err
	}

	if len(rows) == 1 {
		return rows[0], nil
	}
	return nil, errors.New("no data found")
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var user types.User
		row := database.DB.QueryRow("SELECT id, name, phone, password FROM users WHERE id = ?", id)
		err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Password)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, nil
}

func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	var user types.User
	hash, _ := utils.HashPassword(params.Args["password"].(string))
	user.Name = params.Args["name"].(string)
	user.Phone = params.Args["phone"].(string)
	user.Password = hash
	user.Role = params.Args["role"].(string)
	user.Status = params.Args["status"].(string)
	user.CreatedAt = utils.GetTimeNow()
	// fmt.Println(user)

	id, err := query.Insert("users", user, database.DB)
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	forms := map[string]interface{}{
		"name":  params.Args["name"],
		"phone": params.Args["phone"],
	}
	forms["table"] = "users"
	forms["id"] = params.Args["id"].(int)
	err := query.Update(forms, database.DB)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("update user failed")
	}
	id, ok := params.Args["id"].(int)
	if ok {
		var user types.User
		row := database.DB.QueryRow("SELECT id, name, phone, role FROM users WHERE id = ?", id)
		err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Role)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, nil
}

func DeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		stmt, err := database.DB.Prepare("DELETE FROM users WHERE id = ?")
		if err != nil {
			return nil, err
		}
		_, err = stmt.Exec(id)
		if err != nil {
			return nil, err
		}

		return id, nil
	}

	return nil, nil
}
