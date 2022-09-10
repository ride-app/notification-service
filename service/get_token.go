package service

import (
	"context"
	"strings"

	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *NotificationServiceServer) GetToken(ctx context.Context, req *pb.GetNotificationTokenRequest) (*pb.GetNotificationTokenResponse, error)  {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//TODO: implement authentication

	uid := strings.Split(req.Parent, "/")[1]

	token, err := service.tokenRepository.GetToken(ctx, uid)

	if err != nil {
		return nil, status.Error(codes.NotFound, "Token not found")
	}

	return &pb.GetNotificationTokenResponse{
		Token: token,
	}, nil
}