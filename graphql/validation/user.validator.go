package validation

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
)

type ValidationError struct {
	Errors []ValidationErrorItem `json:"errors"`
}

type ValidationErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	var s string
	for _, err := range e.Errors {
		s += fmt.Sprintf("%s: %s\n", err.Field, err.Message)
	}
	return s
}

func ValidateUser(ctx context.Context, user *model.User) (bool, []validation.ValidationError) {
    var validationErrs []validation.ValidationError

    if user.Name == "" {
        validationErrs = append(validationErrs, validation.ValidationError{
            Field:   "name",
            Message: "Name is required",
        })
    }

    if user.Phone == "" {
        validationErrs = append(validationErrs, validation.ValidationError{
            Field:   "phone",
            Message: "Phone is required",
        })
    } else if !regexp.MustCompile(`^[0-9]+$`).MatchString(user.Phone) {
        validationErrs = append(validationErrs, validation.ValidationError{
            Field:   "phone",
            Message: "Phone must be numeric",
        })
    }

    if user.Password == "" {
        validationErrs = append(validationErrs, validation.ValidationError{
            Field:   "password",
            Message: "Password is required",
        })
    } else if len(user.Password) < 6 {
        validationErrs = append(validationErrs, validation.ValidationError{
            Field:   "password",
            Message: "Password must be at least 6 characters long",
        })
    }

    if len(validationErrs) > 0 {
        return false, validationErrs
    }

    return true, nil
}
