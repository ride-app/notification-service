package thirdparty

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/ride-app/notification-service/config"
)

func NewFirebaseApp() (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: config.Env.Firebase_Project_Id}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	app.Database(ctx)

	return app, nil
}
