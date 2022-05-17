package controllers

import (
	"context"
	"farm/src/models"
	pb_log "farm/src/proto/messages/watering_log"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f FarmServer) GetWateringLog (ctx context.Context, request *pb_log.GetWateringLogListRequest) (*pb_log.GetWateringLogListResponse, error) {
	log.Info("Receive message to send watering log specified by time :", ctx.Value("username"))

	result, Error := models.GetWateringLog(request.GetDeviceSerial(), request.GetBeginTime().AsTime(), request.GetEndTime().AsTime())

	var wateringLogSlice []*pb_log.WateringLog
	for _, row := range result {
		var deviceWateringLog pb_log.WateringLog
		deviceWateringLog = pb_log.WateringLog {
			Time:     timestamppb.New(row.LogTime),
			DeviceSerial: row.DeviceSerial,
			WaterAmount: row.WaterAmount,
		}
		wateringLogSlice = append(wateringLogSlice, &deviceWateringLog)
	}

	return &pb_log.GetWateringLogListResponse{
		WateringLog: wateringLogSlice,
	}, Error
}
