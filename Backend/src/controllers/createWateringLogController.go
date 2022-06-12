package controllers

import (
	"context"
	"farm/src/models"
	pb_log "farm/src/proto/messages/watering_log"
	log "github.com/sirupsen/logrus"
	"time"
)

func (f FarmServer) CreateWateringLog(ctx context.Context, request *pb_log.CreateWateringLogRequest) (*pb_log.CreateWateringLogResponse, error) {
	log.Info("Receive message to create watering log: ", ctx.Value("username"))

	device, err := models.GetDevice(request.GetDeviceSerial())
	if err != nil {
		return &pb_log.CreateWateringLogResponse{
			Status: false,
			ErrorMessage: "Device does not exist",
		}, nil
	}
	//time := request.GetTime().AsTime()
	status, errMessage := models.CreateWateringLog(time.Now(), device, request.GetWaterAmount())
	return &pb_log.CreateWateringLogResponse{Status: status, ErrorMessage: errMessage}, nil
}
