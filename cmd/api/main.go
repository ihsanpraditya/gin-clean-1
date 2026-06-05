package main

import (
	"my-api-boilerplate/config"
	"my-api-boilerplate/routes"
)

func main() {
	// Initialize database
	config.ConnectDatabase()

	// Setup routes
	r := routes.SetupRouter()

	// Start server on container port 8080
	r.Run(":8080")
}
