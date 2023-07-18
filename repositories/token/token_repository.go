//go:generate go run github.com/golang/mock/mockgen -destination ../../mocks/$GOFILE -package mocks . TokenRepository

package tokenrepository

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
	"github.com/sirupsen/logrus"
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
		logrus.Info("failed to initialize rtdb")
		logrus.Error(err)
		return nil, err
	}

	return &RTDBImpl{rtdb: rtdb}, nil
}

func (impl *RTDBImpl) GetToken(ctx context.Context, uid string) (*string, error) {
	var token *string
	var data *map[string]interface{}

	logrus.Info("getting token from rtdb")

	if err := impl.rtdb.NewRef(fmt.Sprintf("messaging_tokens/%s", uid)).Get(ctx, data); err != nil {
		logrus.Info("failed to get token from rtdb")
		logrus.Error(err)
		return nil, err
	}

	logrus.Debug(data)

	if token == nil || *token == "" {
		logrus.Info("token not found")
		return nil, nil
	}

	return token, nil
}

func (impl *RTDBImpl) UpdateToken(ctx context.Context, uid string, token string) error {
	logrus.Info("updating token in rtdb")
	err := impl.rtdb.NewRef(fmt.Sprintf("messaging_tokens/%s", uid)).Set(ctx, token)

	if err != nil {
		logrus.Info("failed to update token in rtdb")
		logrus.Error(err)
		return err
	}

	return nil
}
