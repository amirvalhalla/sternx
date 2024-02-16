package main

import (
	"log"
	"time"

	"sternx/config"
	"sternx/infrastructure/logger"
	"sternx/infrastructure/repository"
	"sternx/www/grpc"
)

func main() {
	cfg := config.Default()
	logger.InitGlobalLogger(cfg)

	uow := repository.NewUnitOfWork()
	if err := uow.Setup(cfg.DSN); err != nil {
		logger.Panic("could not setup unit of work and repository", "error", err.Error())
	}

	s := grpc.NewServer(cfg, uow)
	err := s.StartServer()

	log.Println(err)
	for {
		<-time.After(time.Second)
	}
}
