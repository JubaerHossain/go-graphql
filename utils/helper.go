package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CreateJwtToken(id int, role string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
	}).SignedString([]byte("secret"))
}

func ToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

func ToInt(data interface{}) int {
	return int(data.(float64))
}

type ResolverError struct {
	Key     string `json:"key,omitempty"`
	Message string `json:"message"`
}

type ReturnResponse struct {
	Data   interface{}     `json:"data,omitempty"`
	Errors []ResolverError `json:"errors,omitempty"`
	Code   int             `json:"code,omitempty"`
	Status string          `json:"status,omitempty"`
}

func (r *ReturnResponse) Error() string {
	errStr := "Errors: "
	for _, err := range r.Errors {
		errStr += fmt.Sprintf("%s: %s, ", err.Key, err.Message)
	}
	return fmt.Sprintf("%s Code: %d, Status: %s", errStr, r.Code, r.Status)
}

func CreateReturnResponse(data interface{}, errors []ResolverError, code int, status string) *ReturnResponse {
	return &ReturnResponse{
		Data:   data,
		Errors: errors,
		Code:   code,
		Status: status,
	}
}
