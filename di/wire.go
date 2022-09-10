//go:build wireinject

package di

import (
	"github.com/google/wire"
	tokenrepository "github.com/ride-app/notification-service/repositories/token"
	"github.com/ride-app/notification-service/service"
	thirdparty "github.com/ride-app/notification-service/third-party"
)

func InitializeService() (*service.NotificationServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewFirebaseApp,
			tokenrepository.NewRTDBTokenRepository,
			wire.Bind(
				new(tokenrepository.TokenRepository),
				new(*tokenrepository.RTDBImpl),
			),
			service.New,
		),
	)
}
