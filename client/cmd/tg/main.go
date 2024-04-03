package main

import (
	"flag"
	"log"
	"tg_client/internal/config"
	"tg_client/internal/telegram"
)

var fileAddr *string

func init() {
	fileAddr = flag.String("path", "", "enter filename")
}

func main() {
	flag.Parse()
	conf, err := config.InitConfigs(*fileAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err := telegram.StartBot(*conf); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(conf.Token)
}
