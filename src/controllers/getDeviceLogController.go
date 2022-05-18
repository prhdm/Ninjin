package controllers

import (
	"context"
	"farm/src/models"
	pb_device_log "farm/src/proto/messages/device_log"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f FarmServer) GetDeviceLog(ctx context.Context, request *pb_device_log.GetDeviceLogRequest) (*pb_device_log.GetDeviceLogResponse, error) {
	log.Info("Receive message to send log specified by time ")

	result, Error := models.GetDeviceLog(request.GetDeviceSerial(), request.GetBeginTime().AsTime(), request.GetEndTime().AsTime())

	var deviceLogSlice []*pb_device_log.DeviceLog
	for i, row := range result {
		var deviceLog pb_device_log.DeviceLog
		deviceLog = pb_device_log.DeviceLog{
			DeviceSerial: row.DeviceSerial,
			Datetime:     timestamppb.New(row.ServerTime),
			Humidity:     int32(row.Humidity),
			Temp:         int32(row.Temp),
		}
		if i % 10 == 0{
			deviceLogSlice = append(deviceLogSlice, &deviceLog)
		}
	}

	return &pb_device_log.GetDeviceLogResponse{
		DeviceLog: deviceLogSlice,
	}, Error
}
