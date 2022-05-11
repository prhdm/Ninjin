package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func CreateWateringLog(time time.Time ,device *Device, waterAmount float32) (bool, string) {
	wateringLog := WateringLog{LogTime: time, Device: *device, DeviceSerial: device.DeviceSerial, WaterAmount: waterAmount}
	createLog := PostgresDBProvider.DB.Create(&wateringLog)
	if createLog.Error != nil {
		log.Info(createLog.Error)
		return false, "Failed to create watering log in database"
	}
	return true, ""
}
