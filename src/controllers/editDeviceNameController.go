package controllers

import (
	"context"
	"farm/src/models"
	pb_device "farm/src/proto/messages/device"
	log "github.com/sirupsen/logrus"
)
func (f FarmServer) EditDeviceName(ctx context.Context, request *pb_device.EditDeviceNameRequest) (*pb_device.EditDeviceNameResponse, error) {
	log.Info("Receive message to edit device name: ", ctx.Value("username"))

	_, err := models.GetDevice(request.GetDeviceSerial())
	if err != nil {
		return &pb_device.EditDeviceNameResponse{Status: false, ErrorMessage: "Device does not exist"}, nil
	}

	status, errMessage := models.EditDeviceName(request.GetDeviceSerial(), request.GetNewName())
	return &pb_device.EditDeviceNameResponse{Status: status, ErrorMessage: errMessage}, nil
}