package models

import (
	"github.com/jinzhu/gorm"
)

type Farm struct {
	gorm.Model
	FarmName string `gorm:"column:farm_name;"`
}

func (Farm) TableName() string {
	return "farm"
}
