package models

func GetDeviceList() ([]*Device, error) {
	var deviceListSlice []*Device
	result := PostgresDBProvider.DB.Find(&deviceListSlice)
	return deviceListSlice, result.Error
}
