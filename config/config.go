package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewConfig() *Config {
	// Get the absolute path to the directory containing the Go source file
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// Provide the path to the .env file relative to the directory containing the Go source file
	envPath := filepath.Join(dir, "../../../../../../Desktop/go-todo/.env")

	err = godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}

	// Create a new Config struct and populate it with the environment variables
	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	return config
}
