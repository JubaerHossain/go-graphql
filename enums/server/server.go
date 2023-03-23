package server

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"

	"lms/database"
	"lms/graphql/schema"
)

func Start() {

	// database connection
	database.Connect()

	h := handler.New(&handler.Config{
		Schema:     &schema.combinedSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: false,
	})

	http.Handle("/graphql", h)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on port 8081")
	fmt.Println("Access GraphiQL at http://localhost:8081")

	http.ListenAndServe(":8081", nil)
}
