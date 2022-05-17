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

func GetWateringLog(DeviceSerial string, BeginTime time.Time, EndTime time.Time) ([]*WateringLog, error){
	var wateringLogSlice []*WateringLog

	result := PostgresDBProvider.DB.Where("device_serial = ? AND time >= ? AND time <= ?", DeviceSerial, BeginTime, EndTime).Find(&wateringLogSlice)

	return wateringLogSlice, result.Error
}
