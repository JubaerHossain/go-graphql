package resolver

import (
	"errors"
	"fmt"
	"lms/database"
	"lms/graphql/validation"
	"lms/model"
	"lms/utils"
	"reflect"
	"lms/gosql"
	"github.com/graphql-go/graphql"
)

func GetUsers(params graphql.ResolveParams) (interface{}, error) {
	users, err := gosql.QueryModel(reflect.TypeOf(model.User{}), "users", params, database.DB)
	if err != nil {
		return nil, errors.New("no data found")
	}

	return users, nil
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	user, err := gosql.FindByID(reflect.TypeOf(model.User{}), "users", params, database.DB)
	if err != nil {
		return nil, errors.New("no data found")
	}

	return user, nil
}

func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	hash, _ := utils.HashPassword(params.Args["password"].(string))
	userInput := model.User{
		Name:      params.Args["name"].(string),
		Phone:     params.Args["phone"].(string),
		Password:  params.Args["password"].(string),
		Role:      params.Args["role"].(string),
		Status:    params.Args["status"].(string),
		CreatedAt: utils.GetTimeNow(),
	}
	validationErrors := validation.ValidateUser(userInput)
	if validationErrors != nil {
		var errorMsgs []string
		for _, validationErr := range validationErrors {
			errorMsgs = append(errorMsgs, validationErr.Field+" : "+validationErr.Message)
		}
		return nil, fmt.Errorf("%s", errorMsgs)
	}
	userInput.Password = hash
	userInputMap := utils.StructToMap(userInput)
	user, err := gosql.CreateModel(reflect.TypeOf(model.User{}), "users", graphql.ResolveParams{
		Args: map[string]interface{}{
			"model": userInputMap,
		},
	}, database.DB)

	
	fmt.Println(reflect.TypeOf(user))

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	userInput := model.User{
		Name:   params.Args["name"].(string),
		Phone:  params.Args["phone"].(string),
		Role:   params.Args["role"].(string),
		Status: params.Args["status"].(string),
	}
	validationErrors := validation.ValidateUser(userInput)
	if validationErrors != nil {
		var errorMsgs []string
		for _, validationErr := range validationErrors {
			errorMsgs = append(errorMsgs, validationErr.Field+" : "+validationErr.Message)
		}
		return nil, fmt.Errorf("%s", errorMsgs)
	}
	userInputMap := utils.StructToMap(userInput)
	user, err := gosql.UpdateModel(reflect.TypeOf(model.User{}), "users", graphql.ResolveParams{
		Args: map[string]interface{}{
			"model": userInputMap,
		},
	}, database.DB)

	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
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
