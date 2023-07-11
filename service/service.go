package service

import (
	repository "github.com/ride-app/notification-service/repositories/token"
)

type NotificationServiceServer struct {
	tokenRepository repository.TokenRepository
}

func New(
	tokenRepository repository.TokenRepository,
) *NotificationServiceServer {
	return &NotificationServiceServer{
		tokenRepository: tokenRepository,
	}
}
