package service

import (
	"context"
	"strings"

	pb "github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *NotificationServiceServer) UpdateToken(ctx context.Context, req *pb.UpdateNotificationTokenRequest) (*pb.UpdateNotificationTokenResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	uid := strings.Split(req.Name, "/")[1]

	err := service.tokenRepository.UpdateToken(ctx, uid, req.Token)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Inavlid parent")
	}

	return &pb.UpdateNotificationTokenResponse{}, nil
}
