package controllers

import (
	"context"
	"farm/src/models"
	pb_device_list "farm/src/proto/messages/device_list"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) GetDeviceList(ctx context.Context, request *pb_device_list.GetDeviceListRequest) (*pb_device_list.GetDeviceListResponse, error) {
	log.Info("Receive message to get devicesList.")

	devices, err := models.GetDeviceList()
	if err != nil {
		log.Info(err)
	}

	return &pb_device_list.GetDeviceListResponse{
		Devices: parseDeviceQuery(devices),
	}, err
}

func parseDeviceQuery(devices []*models.Device) []*pb_device_list.Device {
	var deviceList []*pb_device_list.Device
	for _, device := range devices {
		deviceList = append(deviceList, &pb_device_list.Device{
			Phone:        device.Phone,
			FarmId:       uint32(device.FarmID),
			DeviceName:   device.DeviceName,
			DeviceSerial: device.DeviceSerial,
		})
	}
	return deviceList
}
