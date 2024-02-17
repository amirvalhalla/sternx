package main

import (
	"sternx/config"
	"sternx/infrastructure/logger"
	"sternx/infrastructure/repository"
	"sternx/www/grpc"
)

func main() {
	cfg, err := config.Default()
	if err != nil {
		logger.Panic("could not initialize config", "error", err.Error())
	}

	logger.InitGlobalLogger(cfg)

	uow := repository.NewUnitOfWork()
	if err := uow.Setup(cfg.DSN); err != nil {
		logger.Panic("could not setup unit of work and repository", "error", err.Error())
	}

	s := grpc.NewServer(cfg, uow)
	err = s.StartServer()
	if err != nil {
		logger.Panic("could not start the grp server", "error", err.Error())
	}

	done := make(chan bool)
	<-done
}
