package controllers

import (
	"farm/src/models"
	pb_warning "farm/src/proto/messages/warning"
	"farm/src/proto/service/farm"
	"time"
)

func (f FarmServer) Warning(request *pb_warning.WarningRequest, stream farm.Farm_WarningServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			warningLogSlice, err := models.GetLastWarnings(uint(request.GetFarmId()))
			if err != nil {
				return err
			}
			for _, warningLog := range warningLogSlice {
				err = stream.SendMsg(&pb_warning.WarningResponse{
					DeviceSerial: warningLog.DeviceSerial,
					Difference: warningLog.Difference,
				})
				if err != nil {
					return err
				}
			}
			time.Sleep(time.Minute * 5)
		}
	}
}
