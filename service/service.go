package service

import (
	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"

	repository "github.com/ride-app/notification-service/repositories/token"
)

type NotificationServiceServer struct {
	pb.UnimplementedNotificationServiceServer

	tokenRepository repository.TokenRepository
}

func New(
	tokenRepository repository.TokenRepository,
) *NotificationServiceServer {
	return &NotificationServiceServer{
		tokenRepository: tokenRepository,
	}
}
