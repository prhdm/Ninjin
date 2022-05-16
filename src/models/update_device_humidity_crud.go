package models

func UpdateDeviceHumidity(deviceSerial string, minHumidity float32, maxHumidity float32) error {
	device := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).First(&Warning{})
	if device.RecordNotFound() {
		return device.Error
	} else {
		err := PostgresDBProvider.DB.Model(&Warning{}).Where("device_serial = ?", deviceSerial).Updates(Warning{
			DeviceSerial: deviceSerial,
			MinHumidity:  minHumidity,
			MaxHumidity:  maxHumidity,
		}).Error
		return err
	}
}
