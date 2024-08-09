package main

import (
	"task_manager/router"
)

// Main entry point of th application where the router is run
func main() {
	r := router.SetupRouter()
	r.Run("localhost:8080")
}
