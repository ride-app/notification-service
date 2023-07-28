package service

import (
	tokenrepository "github.com/ride-app/notification-service/repositories/token"
)

type NotificationServiceServer struct {
	tokenRepository tokenrepository.TokenRepository
}

func New(
	tokenRepository tokenrepository.TokenRepository,
) *NotificationServiceServer {
	return &NotificationServiceServer{
		tokenRepository: tokenRepository,
	}
}
