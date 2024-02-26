package apihandlers

import (
	"github.com/dragonfish/go/v2/pkg/logger"
	tokenrepository "github.com/ride-app/notification-service/internal/repositories/token"
)

type NotificationServiceServer struct {
	tokenRepository tokenrepository.TokenRepository
	logger          logger.Logger
}

func New(
	tokenRepository tokenrepository.TokenRepository,
	logger logger.Logger,
) *NotificationServiceServer {
	return &NotificationServiceServer{
		tokenRepository: tokenRepository,
		logger:          logger,
	}
}
