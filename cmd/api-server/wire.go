//go:build wireinject

package main

import (
	"github.com/deb-tech-n-sol/go/pkg/logger"
	"github.com/google/wire"
	"github.com/ride-app/notification-service/config"
	apihandlers "github.com/ride-app/notification-service/internal/api-handlers"
	tokenrepository "github.com/ride-app/notification-service/internal/repositories/token"
	thirdparty "github.com/ride-app/notification-service/third-party"
)

func InitializeService(logger logger.Logger, config *config.Config) (*apihandlers.NotificationServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewFirebaseApp,
			tokenrepository.NewRTDBTokenRepository,
			wire.Bind(
				new(tokenrepository.TokenRepository),
				new(*tokenrepository.RTDBImpl),
			),
			apihandlers.New,
		),
	)
}
