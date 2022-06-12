package models

import (
	"time"
)

type Warning struct {
	DeviceSerial string  `gorm:"column:device_serial;unique_index;"`
	MinHumidity  float32 `gorm:"column:min_humidity"`
	MaxHumidity  float32 `gorm:"column:max_humidity"`
}

type WarningLog struct {
	DeviceSerial string    `gorm:"column:device_serial"`
	DeviceTime   time.Time `gorm:"column:datetime;" sql:"index:device_time_idx"`
	ServerTime   time.Time `gorm:"column:server_time"`
	Humidity     float32   `gorm:"column:humidity"`
	Difference   float32   `gorm:"column:difference"`
}

func (Warning) TableName() string {
	return "warning"
}
func (WarningLog) TableName() string {
	return "warning_log"
}
