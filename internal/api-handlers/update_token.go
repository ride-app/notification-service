package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/notification-service/api/ride/notification/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *NotificationServiceServer) UpdateNotificationToken(ctx context.Context, req *connect.Request[pb.UpdateNotificationTokenRequest]) (*connect.Response[pb.UpdateNotificationTokenResponse], error) {

	log := service.logger.WithFields(map[string]string{
		"method": "UpdateNotificationToken",
	})

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

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

	err = service.tokenRepository.UpdateToken(ctx, uid, req.Msg.Token)

	if err != nil {
		log.WithError(err).Error("Failed to update token")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&pb.UpdateNotificationTokenResponse{})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully updated token")

	return res, nil
}
