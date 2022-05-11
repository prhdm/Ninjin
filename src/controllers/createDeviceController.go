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
		DeviceSerial: request.DeviceSerial,
		DeviceName:   request.DeviceName,
		Phone:        request.Phone,
		FarmID: uint(request.FarmId),
	}
	errMessage := models.CreateDevice(device)
	var status bool
	var message string
	if errMessage == nil{
		status = true
	} else {
		status = false
		message = errMessage.Error()
	}
	return &pb_device.CreateDeviceResponse{
		Status:       status,
		ErrorMessage: message,
	}, nil
}