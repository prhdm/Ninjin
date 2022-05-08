package models

import(
	"time"
)

func GetDeviceLog(DeviceSerial string, BeginTime time.Time, EndTime time.Time) ([]*DeviceLog, error){
	var deviceLogSlice []*DeviceLog

	result := PostgresDBProvider.DB.Where("device_serial = ? AND datetime >= ? AND datetime <= ?", DeviceSerial, BeginTime, EndTime).Find(&deviceLogSlice)

	return deviceLogSlice, result.Error
}