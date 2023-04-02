package validation

import (
	"context"
	"lms/model"
	"regexp"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateUser(ctx context.Context, user *model.User) (bool, []ValidationErrorItem) {
	var validationErrs []ValidationErrorItem

	if user.Name == "" {
		v := ValidationError{
			Field:   "name",
			Message: "Name is required",
		}
		validationErrs = append(validationErrs, ValidationErrorItem(v))
	}

	if user.Phone == "" {
		v := ValidationError{
			Field:   "phone",
			Message: "Phone is required",
		}
		validationErrs = append(validationErrs, ValidationErrorItem(v))
	} else if !regexp.MustCompile(`^[0-9]+$`).MatchString(user.Phone) {
		v := ValidationError{
			Field:   "phone",
			Message: "Phone must be numeric",
		}
		validationErrs = append(validationErrs, ValidationErrorItem(v))
	}

	if user.Password == "" {
		v := ValidationError{
			Field:   "password",
			Message: "Password is required",
		}
		validationErrs = append(validationErrs, ValidationErrorItem(v))
	} else if len(user.Password) < 6 {
		v := ValidationError{
			Field:   "password",
			Message: "Password must be at least 6 characters long",
		}
		validationErrs = append(validationErrs, ValidationErrorItem(v))
	}

	if len(validationErrs) > 0 {
		return false, validationErrs
	}

	return true, nil
}
