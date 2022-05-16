package models

func SetDeviceHumidity(deviceSerial string, minHumidity float32, maxHumidity float32) error {
	device := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).First(&Device{})
	if device.RecordNotFound() {
		return device.Error
	} else {
		err := PostgresDBProvider.DB.Model(&Warning{}).Create(Warning{
			DeviceSerial: deviceSerial,
			MinHumidity:  minHumidity,
			MaxHumidity:  maxHumidity,
		}).Error
		return err
	}
}
