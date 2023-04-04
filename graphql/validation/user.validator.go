package validation

import (
	"lms/model"

	"github.com/graphql-go/graphql"
)

func ValidateUser(params graphql.ResolveParams) []ValidationErrorItem {
	var user model.User
	user.Name = params.Args["name"].(string)
	user.Phone = params.Args["phone"].(string)
	user.Password = params.Args["password"].(string)
	user.Role = params.Args["role"].(string)
	user.Status = params.Args["status"].(string)
	user.CreatedAt = params.Args["CreatedAt"].(string)
	rules := []ValidationRule{
		{
			Field:       "name",
			Description: "Name",
			Validations: []func(interface{}) ValidationErrorItem{
				RequiredValidation,
				MinLengthValidation(3),
				MaxLengthValidation(50),
			},
		},
		{
			Field:       "phone",
			Description: "Phone",
			Validations: []func(interface{}) ValidationErrorItem{
				RequiredValidation,
				PhoneValidation,
			},
		},
		{
			Field:       "password",
			Description: "Password",
			Validations: []func(interface{}) ValidationErrorItem{
				RequiredValidation,
				MinLengthValidation(6),
			},
		},
		{
			Field:       "role",
			Description: "Role",
			Validations: []func(interface{}) ValidationErrorItem{
				RequiredValidation,
			},
		},
		{
			Field:       "status",
			Description: "Status",
			Validations: []func(interface{}) ValidationErrorItem{
				RequiredValidation,
			},
		},
	}

	return Validate(user, rules)
}
