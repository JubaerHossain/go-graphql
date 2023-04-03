package validation

import (
	"lms/model"
)

func ValidateUser(user model.User) []ValidationErrorItem {
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
