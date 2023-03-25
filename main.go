package main

import (
	"lms/database"
	"lms/server"
)

func main() {
	// database connection
	database.Connect()
	// server start
	server.Start()
}
