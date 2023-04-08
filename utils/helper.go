package utils

import (
	"fmt"
	"reflect"
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
