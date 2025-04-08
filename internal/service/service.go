package service

import (
	"context"
	"fmt"
	"grpc/account/config"

	"github.com/convitnhodev/common/grpc"
	"github.com/convitnhodev/common/logging"
	"github.com/google/wire"
	"go.uber.org/zap"

	_grpc "google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(
	NewService,
)

type Service struct {
	AccountService *accountService
	cfg            *config.Config
	gs             *grpc.Server
	logger         logging.Logger
}

func NewService(ctx context.Context, cfg *config.Config, logger logging.Logger) *Service {
	service := &Service{
		AccountService: NewAccountService(logger),
		cfg:            cfg,
		logger:         logger,
	}

	go func() {
		if err := service.Start(ctx); err != nil {
			logger.Error("Failed to start service", zap.Error(err))
		}
	}()
	return service
}

type hasId interface {
	GetId() string
}

func testIntercepter(ctx context.Context, req interface{}, info *_grpc.UnaryServerInfo, handler _grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("testIntercepter")
	if h, ok := req.(hasId); ok {
		fmt.Println(h.GetId())
	}
	return handler(ctx, req)
}

func (s *Service) Start(ctx context.Context) error {
	gs := grpc.New(
		grpc.WithConfig(s.cfg.GrpcConfig),
		grpc.DisableDefaultLogInterceptor(),
		grpc.WithServiceRegistrar(s.AccountService),
		grpc.WithUnaryServerInterceptor(testIntercepter),
	)

	fmt.Println("Start ok")

	// Store the server in the service
	s.gs = gs

	return gs.Start()
}

func (s *Service) Logger() logging.Logger {
	return s.logger
}
