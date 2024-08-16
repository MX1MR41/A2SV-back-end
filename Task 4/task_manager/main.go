package main

import (
	"task_manager/router"
)

// Main entry point for the application
func main() {
	r := router.SetupRouter() // Setup the router
	r.Run("localhost:8080")   // Run the application at localhost:8080
}
