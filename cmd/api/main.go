package main

import (
	"Demo/router"
)

func main() {
	r := router.SetupRouter()
	r.Run("localhost:8080")
}
