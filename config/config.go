package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	URL      string
	DB       string
	Timeout  int
	User     string
	Password string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func NewConfig(route string) (*Config, error) {
	err := godotenv.Load(route)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := DatabaseConfig{
		URL:      os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DB:       os.Getenv("DB_NAME"),
		Timeout:  30,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}

	serverConfig := ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
		Host: os.Getenv("HOST"),
	}

	config := &Config{
		Database: dbConfig,
		Server:   serverConfig,
	}

	return config, nil
}
