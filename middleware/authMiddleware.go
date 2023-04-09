package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		tokenString := r.Header.Get("Authorization")

		fmt.Println(tokenString)
		fmt.Println("tokenString")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key
			return []byte("secret"), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Set the user ID in the request context
		ctx := context.WithValue(r.Context(), "userID", token.Claims.(jwt.MapClaims)["email"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
