package main

import (
	"Demo/config"
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

	r := router.SetupRouter(db)
	r.Run("localhost:8080")
}
