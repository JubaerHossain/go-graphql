package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	"user-services/database"
	"user-services/graphql/schema"
)

func main() {

	// database connection
	database.Connect()

	h := handler.New(&handler.Config{
		Schema:   &schema.UsersSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Server is running on port 8080")
	fmt.Println("Access GraphiQL at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
