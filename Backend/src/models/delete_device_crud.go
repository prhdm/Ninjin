package models

func DeleteDevice(deviceSerial string) error {
	device := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).First(&Device{})
	if device.RecordNotFound() {
		return device.Error
	} else {
		err := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).Unscoped().Delete(&Device{}).Error
		return err
	}
}
