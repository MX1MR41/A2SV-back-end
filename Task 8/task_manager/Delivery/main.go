package main

import (
	"task_manager/Delivery/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run("localhost:8080")
}
