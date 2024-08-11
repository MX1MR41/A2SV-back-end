package main

import (
	"task_manager/Delivery/routers"
)

// Main entry point of th application where the router is run
func main() {
	r := routers.SetupRouter()
	r.Run("localhost:8080")
}
