package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *NotificationServiceServer) UpdateNotificationToken(ctx context.Context, req *connect.Request[pb.UpdateNotificationTokenRequest]) (*connect.Response[pb.UpdateNotificationTokenResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]
	log.Debug("uid: ", uid)

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	if req.Msg.Token == "" {
		log.Info("Token cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Token cannot be empty")
	}

	err := service.tokenRepository.UpdateToken(ctx, uid, req.Msg.Token)

	if err != nil {
		log.WithError(err).Error("Failed to update token")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully updated token")

	return connect.NewResponse(&pb.UpdateNotificationTokenResponse{}), nil
}
