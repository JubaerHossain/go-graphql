package main

import (
	"fmt"
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
		GraphiQL: false,
		Playground: false,

	})

	http.Handle("/graphql", h)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on port 8080")
	fmt.Println("Access GraphiQL at http://localhost:8080")

	http.ListenAndServe(":8080", nil)


}
