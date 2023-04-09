package resolver

import (
	"errors"
	"lms/database"
	"lms/gosql"
	"reflect"
	"lms/model"
	"lms/utils"

	"github.com/graphql-go/graphql"
)

func Login(params graphql.ResolveParams) (interface{}, error) {
	phone := params.Args["phone"].(string)
	password := params.Args["password"].(string)

	// Prepare the WHERE clause
	where := make(map[string]interface{})
	where["phone"] = phone

	// Query the user by phone number
	users, err := gosql.WhereModel(reflect.TypeOf(model.User{}), "users", params, where, database.DB)
	if err != nil {
		return nil, err
	}

	// Check if the user exists
	if reflect.ValueOf(users).Len() == 0 {
		return nil, errors.New("no data found")
	}

	// Get the first user
	user := reflect.ValueOf(users).Index(0).Interface().(model.User)

	// Compare the password
	err = utils.ComparePassword(password, user.Password)
	if err != nil {
		return nil, err
	}

	// Generate the JWT token and refresh token
	token, err := utils.CreateJwtToken(user.Id)
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.GenerateRefreshToken(user.Id)
	if err != nil {
		return nil, err
	}

	// Return the auth object with the token and refresh token
	auth := model.Auth{
		Token:        token,
		RefreshToken: refreshToken,
	}
	return auth, nil
}

