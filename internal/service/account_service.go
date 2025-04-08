package service

import (
	"context"
	"fmt"
	account "grpc/spec/generated/account"

	"github.com/convitnhodev/common/grpc"
	"github.com/convitnhodev/common/logging"

	"go.uber.org/zap"
)

var _ grpc.ServiceRegistrar = &accountService{}
var _ account.AccountServiceServer = &accountService{}

type accountService struct {
	account.UnimplementedAccountServiceServer
	logger logging.Logger
}

func NewAccountService(logger logging.Logger) *accountService {
	return &accountService{
		logger: logger,
	}
}

func (svc *accountService) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	return nil, nil
}

func (svc *accountService) GetAccount(ctx context.Context, req *account.GetAccountRequest) (*account.GetAccountResponse, error) {
	svc.logger.Info("GetAccount", zap.Any("req", req))
	return &account.GetAccountResponse{
		Code:    0,
		Message: "success",
		Data: &account.GetAccountResponse_Data{
			Id:   req.Id,
			Name: "hung dep trai",
		},
	}, nil
}

func (svc *accountService) RegisterService(server *grpc.Server) {
	fmt.Println("RegisterService")
	account.RegisterAccountServiceServer(server, svc)
	fmt.Println("RegisterService ok")
}
