package utils

import (
	"errors"
	"fmt"
	"lms/database"
	"lms/gosql"
	"reflect"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(plainPassword string, hashedPassword string) error {
	// Compare the plaintext password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CreateJwtToken(userID int) (string, error) {
	// Set the JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Set the token expiration time to 24 hours
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the JWT token with a secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID int) (string, error) {
	// Generate a random UUID
	refreshToken := uuid.New().String()

	// Set the expiration time for the refresh token
	expirationTime := time.Now().Add(time.Hour * 24 * 7) // Set the expiration time to 7 days
	// Store the refresh token in the database with the corresponding user ID and expiration time
	data := map[string]interface{}{
		"user_id":         userID,
		"token":           refreshToken,
		"expiration_time": expirationTime.Format("2006-01-02 15:04:05"),
	}

	// Call the RawInsertModel function to insert the data
	_, err := gosql.RawInsertModel("refresh_tokens", data, database.DB)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
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

func StructToMap(obj interface{}) map[string]interface{} {
	v := reflect.ValueOf(obj)
	values := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		values[v.Type().Field(i).Name] = v.Field(i).Interface()
	}
	return values
}

func MapToStruct(data map[string]interface{}, resultType reflect.Type) interface{} {
	result := reflect.New(resultType).Elem()
	for key, value := range data {
		field := result.FieldByName(key)
		if !field.IsValid() {
			continue
		}
		if !field.CanSet() {
			continue
		}
		fieldValue := reflect.ValueOf(value)
		if field.Type() != fieldValue.Type() {
			continue
		}
		field.Set(fieldValue)
	}
	return result.Interface()
}

func SetStructField(s interface{}, name string, value interface{}) error {
	v := reflect.ValueOf(s).Elem()
	f := v.FieldByName(name)

	if !f.IsValid() {
		return fmt.Errorf("no such field: %s in struct", name)
	}

	if !f.CanSet() {
		return fmt.Errorf("cannot set field %s value", name)
	}

	val := reflect.ValueOf(value)
	if f.Type() != val.Type() {
		return fmt.Errorf("value type %s does not match field type %s", val.Type(), f.Type())
	}

	f.Set(val)
	return nil
}
