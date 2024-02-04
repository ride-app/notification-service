package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/notification-service/api/ride/notification/v1alpha1"
)

func (service *NotificationServiceServer) GetNotificationToken(ctx context.Context, req *connect.Request[pb.GetNotificationTokenRequest]) (*connect.Response[pb.GetNotificationTokenResponse], error) {

	log := service.logger.WithFields(map[string]string{
		"method": "GetNotificationToken",
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

	token, err := service.tokenRepository.GetToken(ctx, log, uid)

	if err != nil {
		log.WithError(err).Error("Failed to get token")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if token == nil {
		log.Info("Token not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("token not found"))
	}

	res := connect.NewResponse(&pb.GetNotificationTokenResponse{
		Token: *token,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Debug("token: ", *token)
	log.Info("Successfully retrieved token")

	return res, nil
}
