package main

import (
	"lms/database"
	"lms/enums/server"
)

func main() {
	// database connection
	database.Connect()
	// server start
	server.Start()
}
