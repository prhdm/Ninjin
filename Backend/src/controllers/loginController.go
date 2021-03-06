package controllers

import (
	"context"
	"farm/src/models"
	pb_user "farm/src/proto/messages/user"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) Login(ctx context.Context, request *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	log.Info("Receive message to login: ", request.Username)
	user, err := models.GetUser(request.Username)
	if err != nil {
		return &pb_user.LoginResponse{
			Status: false, ErrorMessage: "Username or password not valid.",
		}, nil
	}
	status, errMessage, token := models.Login(user, request.Password)
	return &pb_user.LoginResponse{
		Status:       status,
		ErrorMessage: errMessage,
		Token:        token,
	}, nil
}
