package models

func GetDeviceList(farmId uint) ([]*Device, error) {
	var deviceListSlice []*Device
	result := PostgresDBProvider.DB.Where("farm_id = ?", farmId).Find(&deviceListSlice)
	return deviceListSlice, result.Error
}
