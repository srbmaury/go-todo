package config

import (
	todoModels "Todos/internal/todos/models"
	userModels "Todos/internal/users/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (c *Config) ConnectDB() (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Errror is", err)
		return nil, err
	}

	err = db.AutoMigrate(&todoModels.Todo{}, &userModels.User{})

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
