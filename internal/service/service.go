package service

import (
	"context"
	"fmt"
	"grpc/account/config"

	"github.com/convitnhodev/common/grpc"
	"github.com/google/wire"

	_grpc "google.golang.org/grpc"

	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewService,
	wire.Bind(new(GrpcServer), new(*Service)),
)

type Service struct {
	AccountService *accountService
	cfg            *config.Config
	gs             *grpc.Server
	logger         *zap.Logger
}

type GrpcServer interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Logger() *zap.Logger
}

func NewService(ctx context.Context, cfg *config.Config, logger *zap.Logger) *Service {
	service := &Service{
		AccountService: NewAccountService(logger),
		cfg:            cfg,
		logger:         logger,
	}

	service.Start(ctx)
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

func (s *Service) Stop(ctx context.Context) error {
	s.gs.Shutdown()
	return nil
}

func (s *Service) Logger() *zap.Logger {
	return s.logger
}
