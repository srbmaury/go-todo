package main

import (
	"github.com/srbmaury/go-todo/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run("localhost:8080")
}
