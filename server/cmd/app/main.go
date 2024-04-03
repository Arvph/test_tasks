package main

import (
	"flag"

	"github.com/arvph/test_tasks/internal/configs"
	"github.com/arvph/test_tasks/internal/database"
	"github.com/arvph/test_tasks/internal/repository"
	"github.com/arvph/test_tasks/internal/server"
	"github.com/arvph/test_tasks/internal/services"
	"github.com/sirupsen/logrus"
)

var fileAddr *string

func init() {
	fileAddr = flag.String("path", "", "enter filename")
}

func main() {
	log := logrus.New()

	flag.Parse()
	// запустить конфиги
	conf, err := configs.InitConfigs(log, *fileAddr)
	if err != nil {
		log.Fatal(err)
	}

	// запустить подключение к БД
	if err := database.PoolConnects(log, &conf.DB); err != nil {
		log.Fatal(err)
	}
	defer conf.DB.ClosePool(log)

	// подключение сервисов
	conf.Server.SetServices(services.NewService(repository.NewRepository(&conf.DB)))

	// запустить сервер
	if err := server.ServerStart(log, &conf.Server); err != nil {
		log.Fatal(err)
	}

}
