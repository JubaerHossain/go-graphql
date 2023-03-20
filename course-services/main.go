package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"

	"course-services/database"
	"course-services/graphql/schema"
)

func main() {

	// database connection
	database.Connect()

	h := handler.New(&handler.Config{
		Schema: &schema.CourseSchema,
		Pretty: true,
		GraphiQL: false,
		Playground: false,
	})


	http.Handle("/graphql", h)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on port 8081")
	fmt.Println("Access GraphiQL at http://localhost:8081")

	http.ListenAndServe(":8081", nil)
}
