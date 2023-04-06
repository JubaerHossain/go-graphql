package validation

import (
	"lms/model"

	"github.com/JubaerHossain/validation"
)

func ValidateUser(user model.User) []validation.ValidationErrorItem {

	rules := []validation.ValidationRule{
		{
			Field:       "name",
			Description: "Name",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.RequiredValidation,
				validation.MinLengthValidation(3),
				validation.MaxLengthValidation(50),
			},
		},
		{
			Field:       "phone",
			Description: "Phone",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.RequiredValidation,
				validation.PhoneValidation,
			},
		},
		{
			Field:       "password",
			Description: "Password",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.RequiredValidation,
				validation.MinLengthValidation(6),
			},
		},
		{
			Field:       "role",
			Description: "Role",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.RequiredValidation,
			},
		},
		{
			Field:       "status",
			Description: "Status",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.RequiredValidation,
			},
		},
	}

	return validation.Validate(user, rules)
}
