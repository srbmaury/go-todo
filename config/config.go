package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     "127.0.0.1",
		DBPort:     "3306",
		DBUser:     "root",
		DBPassword: "hell0@1ntern123",
		DBName:     "todos",
	}
}
