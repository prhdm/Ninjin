package models

func GetDevice(deviceSerial string) (*Device, error) {
	device := &Device{}
	result := PostgresDBProvider.DB.Where("device_serial = ?", deviceSerial).Find(&device)
	return device, result.Error
}
