package controllers

import (
	"context"
	pb_update_device "farm/src/proto/messages/update_device"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) UpdateDevice(ctx context.Context, request *pb_update_device.UpdateDeviceRequest) (*pb_update_device.UpdateDeviceResponse, error) {
	log.Info("Receive message to update devices.")

	return &pb_update_device.UpdateDeviceResponse{
		Status: true,
	}, nil
}
