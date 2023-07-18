//go:generate go run github.com/golang/mock/mockgen -destination ../../mocks/$GOFILE -package mocks . TokenRepository

package tokenrepository

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
	log "github.com/sirupsen/logrus"
)

type TokenRepository interface {
	GetToken(ctx context.Context, uid string) (*string, error)

	UpdateToken(ctx context.Context, uid string, token string) error
}

type RTDBImpl struct {
	rtdb *db.Client
}

func NewRTDBTokenRepository(firebaseApp *firebase.App) (*RTDBImpl, error) {
	rtdb, err := firebaseApp.Database(context.Background())

	if err != nil {
		log.Info("failed to initialize rtdb")
		log.Error(err)
		return nil, err
	}

	return &RTDBImpl{rtdb: rtdb}, nil
}

func (impl *RTDBImpl) GetToken(ctx context.Context, uid string) (*string, error) {
	var token *string = new(string)

	log.Info("getting token from rtdb")
	log.Debug(fmt.Sprintf("path: messaging_tokens/%s", uid))

	if err := impl.rtdb.NewRef(fmt.Sprintf("messaging_tokens/%s", uid)).Get(ctx, token); err != nil {
		log.Errorf("failed to get token from rtdb: %v", err)
		return nil, err
	}

	if token == nil || *token == "" {
		log.Info("token not found")
		return nil, nil
	}

	return token, nil
}

func (impl *RTDBImpl) UpdateToken(ctx context.Context, uid string, token string) error {
	log.Info("updating token in rtdb")

	if err := impl.rtdb.NewRef(fmt.Sprintf("messaging_tokens/%s", uid)).Set(ctx, token); err != nil {
		log.Info("failed to update token in rtdb")
		log.Error(err)
		return err
	}

	return nil
}
