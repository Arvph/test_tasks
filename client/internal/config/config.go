package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func InitConfigs(filename string) (*Config, error) {
	log.Println("Trying to open .env file")

	if err := godotenv.Load(filename); err != nil {
		log.Println("Unable to open .env file", err)
		return nil, err
	}
	log.Println(".env file opened")
	return &Config{
		Token: os.Getenv("TOKEN"),
	}, nil
}
