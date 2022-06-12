package controllers

import (
	"context"
	"farm/src/models"
	pb_set_device_humidity "farm/src/proto/messages/set_device_humidity"

	log "github.com/sirupsen/logrus"
)

func (f FarmServer) SetDeviceHumidity(ctx context.Context, request *pb_set_device_humidity.SetDeviceHumidityRequest) (*pb_set_device_humidity.SetDeviceHumidityResponse, error) {
	log.Info("Receive message to set Device " + request.DeviceSerial + " humidity.")
	err := models.SetDeviceHumidity(request.DeviceSerial, request.MinHumidity, request.MaxHumidity)
	if err != nil {
		return &pb_set_device_humidity.SetDeviceHumidityResponse{
			Status:       err == nil,
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb_set_device_humidity.SetDeviceHumidityResponse{
		Status: err == nil,
	}, nil
}
