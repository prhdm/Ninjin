package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type WateringLog struct {
	gorm.Model
	LogTime time.Time `gorm:"column:time"`
	DeviceSerial   string `gorm:"column:device_serial"`
	Device Device
	WaterAmount float32 `gorm:"column:water_amount"`
}
