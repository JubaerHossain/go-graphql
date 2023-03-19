package resolver

import (
	"database/sql"
	"errors"
	"fmt"
	"user-services/database"
	"user-services/graphql/types"
	"user-services/utils"

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
	fmt.Println(offset)
	rows, err := database.DB.Query("SELECT id, name, phone FROM users ORDER BY id DESC LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var user types.User
		var phone sql.NullString
		err := rows.Scan(&user.ID, &user.Name, &phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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

	id, err := database.INSERT("users", user)
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}


func UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	var user types.User
	fmt.Println(params.Args)
	user.ID = params.Args["id"].(int)
	user.Name = params.Args["name"].(string)
	user.Phone = params.Args["phone"].(string)
	
	update, err := database.UPDATE("users", user, "id=?", user.ID)
	if err != nil {
		return nil, errors.New("update user failed")
        return nil, err
    }

	fmt.Println("update")
	fmt.Println(update)

	updatedUser, err := database.FindByID("users", user.ID)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
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
