package controllers

import (
	"context"
	"farm/src/models"
	pb_device "farm/src/proto/messages/device"
	log "github.com/sirupsen/logrus"
)

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