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
	hash, _ := utils.HashPassword(params.Args["password"].(string))
	params.Args["password"] = hash
	params.Args["role"] = "user"
	params.Args["CreatedAt"] = utils.GetTimeNow()
	validationErrors := validation.ValidateUser(params)
	if validationErrors != nil {
		var errorMsgs []string
		for _, validationErr := range validationErrors {
			errorMsgs = append(errorMsgs, validationErr.Field+" : "+validationErr.Message)
		}
		return nil, fmt.Errorf("%s", errorMsgs)
	}
	user, err := query.CreateModel(reflect.TypeOf(model.User{}), "users", params)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("no data found")
	}

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
