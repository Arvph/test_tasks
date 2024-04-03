package configs

import (
	"os"

	"github.com/arvph/test_tasks/internal/database"
	"github.com/arvph/test_tasks/internal/server"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

// Configs представляет структуру конфигов
type Configs struct {
	Server server.Server
	DB     database.DB
}

// InitConfigs парсит данные из .env
func InitConfigs(log *logrus.Logger, filename string) (*Configs, error) {
	log.Println("Trying to open .env file")

	if err := godotenv.Load(filename); err != nil {
		log.Println("Unable to open .env file", err)
		return nil, err
	}
	log.Println(".env file opened")
	return &Configs{
		Server: server.Server{
			Addr: os.Getenv("S_URL"),
			Port: os.Getenv("S_PORT"),
		},
		DB: database.DB{
			DbPort:     os.Getenv("DB_PORT"),
			DbHost:     os.Getenv("DB_DSN"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}, nil
}
