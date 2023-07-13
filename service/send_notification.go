package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *NotificationServiceServer) SendNotification(ctx context.Context, req *connect.Request[pb.SendNotificationRequest]) (*connect.Response[pb.SendNotificationResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//TODO: implement authentication

	return connect.NewResponse(&pb.SendNotificationResponse{}), nil
}
