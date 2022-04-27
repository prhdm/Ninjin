package controllers

import (
	"context"
	"farm/src/models"
	pb_device "farm/src/proto/messages/device"
	pb_user "farm/src/proto/messages/user"
	log "github.com/sirupsen/logrus"
)

type FarmServer struct {
}

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

func (f FarmServer) CreateDevice(ctx context.Context, request *pb_device.CreateDeviceRequest) (*pb_device.CreateDeviceResponse, error) {
	log.Info("Receive message to CreateDevice")
	device := &models.Device{
		DeviceSerial: request.Device.DeviceSerial,
		Phone:        request.Device.Phone,
	}
	errMessage := models.CreateDevice(device)
	status := false
	if errMessage == nil{
		status = true
	}
	return &pb_device.CreateDeviceResponse{
		Status:       status,
		ErrorMessage: "there was a problem!",
	}, nil
}