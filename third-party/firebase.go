package thirdparty

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish-tech/go/pkg/logger"
	"github.com/ride-app/notification-service/config"
)

func NewFirebaseApp(log logger.Logger, config *config.Config) (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID:   config.ProjectID,
		DatabaseURL: config.FirebaseDatabaseUrl,
	}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.WithError(err).Fatal("Cannot initialize firebase app")
		return nil, err
	}

	return app, nil
}
