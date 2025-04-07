//go:build wireinject
// +build wireinject

package app

import (
	"context"
	config "grpc/account/config"
	"grpc/account/internal/service"

	"github.com/google/wire"
	logging "gitlab.zalopay.vn/fin/platform/common/logging"
)

// Provide context
func provideContext() context.Context {
	return context.Background()
}

// Provide logger config
func provideLoggerConfig(cfg *config.Config) *logging.Config {
	return config.LoggerConfig
}

var AppSet = wire.NewSet(
	provideContext,
	provideLoggerConfig,
)

func InitializeApp(cfg *config.Config) (*service.Service, error) {
	wire.Build(
		// App-level dependencies
		AppSet,
		logging.ProviderSet,
		// Service dependencies
		service.ProviderSet,
	)
	return nil, nil
}
