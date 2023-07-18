package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func (service *NotificationServiceServer) GetNotificationToken(ctx context.Context, req *connect.Request[pb.GetNotificationTokenRequest]) (*connect.Response[pb.GetNotificationTokenResponse], error) {

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

	token, err := service.tokenRepository.GetToken(ctx, uid)

	if err != nil {
		log.Error("Failed to get token: ", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if token == nil {
		log.Info("Token not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("token not found"))
	}

	log.Debug("token: ", *token)
	log.Info("Successfully retrieved token")

	return connect.NewResponse(&pb.GetNotificationTokenResponse{
		Token: *token,
	}), nil
}
