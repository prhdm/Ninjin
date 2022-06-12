package models

import (
	"time"
)

func GetDeviceLog(DeviceSerial string, BeginTime time.Time, EndTime time.Time, farmId uint) ([]*DeviceLog, error){
	var deviceLogSlice []*DeviceLog
	device := PostgresDBProvider.DB.Where("device_serial = ? AND farm_id = ?", DeviceSerial,farmId).First(&Device{})

	if device.RecordNotFound() {
		return deviceLogSlice,device.Error
	}

	result := PostgresDBProvider.DB.Where("device_serial = ? AND server_time >= ? AND server_time <= ?", DeviceSerial, BeginTime, EndTime).Find(&deviceLogSlice)

	return deviceLogSlice, result.Error
}