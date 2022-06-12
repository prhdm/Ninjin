package models

func SetDeviceHumidity(deviceSerial string, minHumidity float32, maxHumidity float32) error {
	device := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).First(&Device{})
	if device.RecordNotFound() {
		return device.Error
	} else {
		device = PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).First(&Warning{})
		if device.RecordNotFound() {
			err := setDeviceHumidityInDatabase(deviceSerial, minHumidity, maxHumidity)
			return err
		} else {
			err := updateDeviceHumidity(deviceSerial, minHumidity, maxHumidity)
			return err
		}
	}
}

func setDeviceHumidityInDatabase(deviceSerial string, minHumidity float32, maxHumidity float32) error {
	err := PostgresDBProvider.DB.Model(&Warning{}).Create(Warning{
		DeviceSerial: deviceSerial,
		MinHumidity:  minHumidity,
		MaxHumidity:  maxHumidity,
	}).Error
	return err
}

func updateDeviceHumidity(deviceSerial string, minHumidity float32, maxHumidity float32) error {
	err := PostgresDBProvider.DB.Model(&Warning{}).Where("device_serial = ?", deviceSerial).Updates(Warning{
		DeviceSerial: deviceSerial,
		MinHumidity:  minHumidity,
		MaxHumidity:  maxHumidity,
	}).Error
	return err
}
