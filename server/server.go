package server

import (
	"fmt"
	"lms/graphql/schema"
	"lms/routes"
	"net/http"

	"lms/database"

	"github.com/graphql-go/handler"
)

func Start() {

	// database connection
	database.Connect()

	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: false,
	})

	http.Handle("/graphql", h)
	http.Handle("/query", http.StripPrefix("/query", http.FileServer(http.Dir("static"))))
	routes.Init()
	fmt.Println("Server is running on port 8081")
	fmt.Println("Access GraphiQL at http://localhost:8081/query")

	http.ListenAndServe(":8081", nil)
}
