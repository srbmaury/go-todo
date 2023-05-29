package main

import (
	"Demo/config"
	context "Demo/context"
	"Demo/router"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	context.CreateContext()
	r := router.SetupRouter(db)
	r.Run("localhost:8080")
}
