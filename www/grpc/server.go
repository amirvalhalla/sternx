package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"sternx/config"
	"sternx/infrastructure/logger"
	"sternx/infrastructure/repository"
	sternx "sternx/www/grpc/gen/go"
)

type Server struct {
	ctx      context.Context
	cancel   context.CancelFunc
	listener net.Listener
	address  string
	grpc     *grpc.Server
	config   *config.Config
	uow      repository.UnitOfWork
	logger   *logger.SubLogger
}

func NewServer(cfg *config.Config, uow repository.UnitOfWork) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		ctx:    ctx,
		cancel: cancel,
		config: cfg,
		uow:    uow,
		logger: logger.NewSubLogger("_grpc", nil),
	}
}

func (s *Server) StartServer() error {
	listener, err := net.Listen("tcp", s.config.GRPCListen)
	if err != nil {
		return err
	}

	return s.startListening(listener)
}

func (s *Server) startListening(listener net.Listener) error {
	grpcServer := grpc.NewServer()

	userGRPCServer := newUserServer(s, s.uow, s.config)

	sternx.RegisterUserServer(grpcServer, userGRPCServer)

	s.listener = listener
	s.address = listener.Addr().String()
	s.grpc = grpcServer

	s.logger.Info("grpc started listening", "address", listener.Addr().String())
	go func() {
		if err := s.grpc.Serve(listener); err != nil {
			s.logger.Error("error on grpc serve", "error", err)
		}
	}()

	return s.startGateway(s.address)
}
