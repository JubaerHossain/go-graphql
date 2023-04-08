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

func ValidateUserLogin(user model.User) []validation.ValidationErrorItem {

	rules := []validation.ValidationRule{
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
	}

	return validation.Validate(user, rules)
}

func ValidateUserUpdate(user model.User) []validation.ValidationErrorItem {

	rules := []validation.ValidationRule{
		{
			Field:       "name",
			Description: "Name",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.MinLengthValidation(3),
				validation.MaxLengthValidation(50),
				validation.StringValidation,
			},
		},
		{
			Field:       "phone",
			Description: "Phone",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.PhoneValidation,
			},
		},
		{
			Field:       "status",
			Description: "Status",
			Validations: []func(interface{}) validation.ValidationErrorItem{
				validation.MinLengthValidation(3),
			},
		},
	}

	return validation.Validate(user, rules)
}
