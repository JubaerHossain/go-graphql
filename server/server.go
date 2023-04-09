package server

import (
	"fmt"
	"lms/graphql/schema"
	"lms/routes"
	"net/http"
	"os"

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
		Playground: true,
	})

	APP_PORT := os.Getenv("APP_PORT")
	APP_URL := os.Getenv("APP_URL")

	http.Handle("/graphql", h)
	routes.Init()
	fmt.Println("Server is running on port " + APP_PORT)
	fmt.Println("Access GraphiQL at " + APP_URL + ":" + APP_PORT + "/graphql")

	err := http.ListenAndServe(":"+APP_PORT, nil)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}

}
