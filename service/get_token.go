package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
)

func (service *NotificationServiceServer) GetNotificationToken(ctx context.Context, req *connect.Request[pb.GetNotificationTokenRequest]) (*connect.Response[pb.GetNotificationTokenResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	//TODO: implement authentication

	uid := strings.Split(req.Msg.Name, "/")[1]

	token, err := service.tokenRepository.GetToken(ctx, uid)

	if token == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Token not found"))
	}

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.GetNotificationTokenResponse{
		Token: *token,
	}), nil
}
