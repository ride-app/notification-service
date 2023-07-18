package thirdparty

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/ride-app/notification-service/config"
	log "github.com/sirupsen/logrus"
)

func NewFirebaseApp() (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID:   config.Env.Firebase_Project_Id,
		DatabaseURL: config.Env.Firebase_Database_url,
	}
	log.Info("initializing firebase app")
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Info("failed to initialize firebase app")
		log.Fatalln(err)
		return nil, err
	}

	return app, nil
}
