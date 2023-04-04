package resolver

import (
	"errors"
	"fmt"
	"lms/database"
	"lms/graphql/validation"
	"lms/model"
	"lms/query"
	"lms/utils"
	"reflect"

	"github.com/graphql-go/graphql"
)

func GetUsers(params graphql.ResolveParams) (interface{}, error) {
	users, err := query.QueryModel(reflect.TypeOf(model.User{}), "users", params)
	if err != nil {
		return nil, errors.New("no data found")
	}

	return users, nil
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	user, err := query.FindByID(reflect.TypeOf(model.User{}), "users", params)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no data found")
	}

	return user, nil
}

func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	var user model.User
	hash, _ := utils.HashPassword(params.Args["password"].(string))
	user.Name = params.Args["name"].(string)
	user.Phone = params.Args["phone"].(string)
	user.Password = hash
	user.Role = params.Args["role"].(string)
	user.Status = params.Args["status"].(string)
	user.CreatedAt = utils.GetTimeNow()
	// fmt.Println(user)

	validationErrors := validation.ValidateUser(user)
	if validationErrors != nil {
		var errorMsgs []string
		for _, validationErr := range validationErrors {
			errorMsgs = append(errorMsgs, validationErr.Field+" : "+validationErr.Message)
		}
		return errorMsgs, fmt.Errorf("validation error")
	}

	params.Args["password"] = hash

	fmt.Println(params.Args)

	// user, err := query.CreateModel(reflect.TypeOf(model.User{}), "users", params)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, errors.New("no data found")
	// }

	return nil, nil

	// continue with creating the user

	id, err := query.Insert("users", user, database.DB)
	if err != nil {
		return nil, err
	}
	user.Id = int(id)
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
		var user model.User
		row := database.DB.QueryRow("SELECT id, name, phone, role FROM users WHERE id = ?", id)
		err := row.Scan(&user.Id, &user.Name)
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
