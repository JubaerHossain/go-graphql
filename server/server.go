package server

import (
	"fmt"
	"lms/graphql/schema"
	"lms/middleware"
	"lms/routes"
	"net/http"
	"os"

	"lms/database"

	"github.com/graphql-go/handler"
	"github.com/rs/cors"
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

	handler := cors.Default().Handler(h)

	http.Handle("/graphql", middleware.Authenticate(handler))
	routes.Init()
	fmt.Println("Server is running on port " + APP_PORT)
	fmt.Println("Access GraphiQL at " + APP_URL + ":" + APP_PORT + "/graphql")

	http.ListenAndServe(":"+APP_PORT, nil)

}
