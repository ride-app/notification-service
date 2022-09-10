package thirdparty

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/ride-app/notification-service/config"
)

func NewFirebaseApp() (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: config.Env.Firebase_Project_Id,
		DatabaseURL: config.Env.Firebase_Database_url,
	}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return app, nil
}
