package main

import (
	"log"
	"time"

	"sternx/config"
	"sternx/infrastructure/logger"
	"sternx/www/grpc"
)

func main() {
	cfg := config.Default()
	logger.InitGlobalLogger(cfg)
	s := grpc.NewServer(cfg)
	err := s.StartServer()

	log.Println(err)
	for {
		<-time.After(time.Second)
	}
}
