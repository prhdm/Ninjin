package controllers

import (
	"context"
	"farm/src/models"
	pb_update_device_humidity "farm/src/proto/messages/update_device_humidity"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) UpdateDeviceHumidity(ctx context.Context, request *pb_update_device_humidity.UpdateDeviceHumidityRequest) (*pb_update_device_humidity.UpdateDeviceHumidityResponse, error) {
	log.Info("Receive message to update Device " + request.DeviceSerial + " humidity.")
	err := models.UpdateDeviceHumidity(request.DeviceSerial, request.MinHumidity, request.MaxHumidity)
	if err != nil {
		return &pb_update_device_humidity.UpdateDeviceHumidityResponse{
			Status:       err == nil,
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb_update_device_humidity.UpdateDeviceHumidityResponse{
		Status: err == nil,
	}, nil
}
