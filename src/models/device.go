package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model
	DeviceSerial string `gorm:"column:device_serial"`
	Phone    string `gorm:"column:phone"`
	Farm     Farm
	FarmID   uint
	DeviceName string `gorm:"column:device_name"`
}

type DeviceLog struct {
	DeviceSerial   string `gorm:"column:device_serial"`
	Device     Device
	DeviceTime time.Time `gorm:"column:datetime;" sql:"index:device_time_idx"`
	ServerTime time.Time `gorm:"column:server_time"`
	Humidity   float32   `gorm:"column:humidity"`
	Temp       float32   `gorm:"column:temp"`
}

func (Device) TableName() string {
	return "device"
}

func (DeviceLog) TableName() string {
	return "device_log"
}
