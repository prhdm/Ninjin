package controllers

import (
	"context"
	"farm/src/models"
	pb_user "farm/src/proto/messages/user"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) Logout(ctx context.Context, request *pb_user.LogoutRequest) (*pb_user.LogoutResponse, error) {
	log.Info("Receive message to logout: ", ctx.Value("username"))

	encodedToken, _ := grpc_auth.AuthFromMD(ctx, "bearer")
	status := models.Logout(encodedToken)
	return &pb_user.LogoutResponse{Status: status}, nil
}