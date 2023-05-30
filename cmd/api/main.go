package main

import (
	"Todos/config"
	context "Todos/context"
	"Todos/router"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	context.CreateContext()
	r := router.SetupRouter(db)
	r.Run("localhost:8080")
}
