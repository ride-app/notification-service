package service

import (
	"context"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *NotificationServiceServer) UpdateNotificationToken(ctx context.Context, req *connect.Request[pb.UpdateNotificationTokenRequest]) (*connect.Response[pb.UpdateNotificationTokenResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]

	if req.Msg.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "Token cannot be empty")
	}

	err := service.tokenRepository.UpdateToken(ctx, uid, req.Msg.Token)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.UpdateNotificationTokenResponse{}), nil
}
