package main

import (
	"OnlyID/config"
	"OnlyID/library/log"
	"flag"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(config.Conf.Log)
	//s := service.NewService(config.Conf)

}
