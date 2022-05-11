package controllers

import (
	"context"
	"farm/src/models"
	pb_delete_device "farm/src/proto/messages/delete_device"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) DeleteDevice(ctx context.Context, request *pb_delete_device.DeleteDeviceRequest) (*pb_delete_device.DeleteDeviceResponse, error) {
	log.Info("Receive message to delete device.")
	err := models.DeleteDevice(request.DeviceSerial)
	return &pb_delete_device.DeleteDeviceResponse{
		Status: err == nil,
	}, err
}
