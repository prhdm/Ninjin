package models

import log "github.com/sirupsen/logrus"

func GetDevice(deviceSerial string) (*Device, error) {
	device := &Device{}
	result := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).Find(&device)
	return device, result.Error
}

func EditDeviceName(deviceSerial string, newDeviceName string) (bool, string) {
	result := PostgresDBProvider.DB.Model(&Device{}).Where("device_serial = ?", deviceSerial).Update("device_name", newDeviceName)
	if result.Error != nil {
		log.Info(result.Error)
		return false, "Failed to update name of the device in database"
	}
	return true, ""
}
