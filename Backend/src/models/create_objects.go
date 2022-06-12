package models

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Create insert the value into database

func CreateFarm(farmname string) (uint, error) {
	farm := &Farm{
		FarmName: farmname,
	}
	if err := PostgresDBProvider.DB.Create(farm).Error; err != nil {
		return 0, err
	}
	return farm.ID, nil
}

func CreateUser(username string, password string, farmID uint) error {
	user := &User{}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("Password Encryption  failed")
		return err
	}
	user.Password = string(pass)
	user.Username = username
	user.FarmID = uint(farmID)
	createdUser := PostgresDBProvider.DB.Create(user)

	if createdUser.Error != nil {
		log.Info(createdUser.Error)
		return createdUser.Error
	}
	return nil
}

func CreateDeviceLog(serial string, humidity int32, deviceTime, serverTime time.Time) error {
	deviceLog := &DeviceLog{
		DeviceSerial: serial,
		DeviceTime:   deviceTime,
		ServerTime:   serverTime,
		Humidity:     float32(humidity),
		Temp:         0,
	}
	err := warningLog(serial, humidity, deviceTime, serverTime)
	if err != nil {
		log.Info(err)
		return err
	}
	createdDataLog := PostgresDBProvider.DB.Create(deviceLog)
	if createdDataLog.Error != nil {
		log.Info(createdDataLog.Error)
		return createdDataLog.Error
	}
	return nil
}

func warningLog(serial string, humidity int32, deviceTime, serverTime time.Time) error {
	var device Warning
	if !PostgresDBProvider.DB.Where("device_serial = ?", serial).First(&device).RecordNotFound() {
		difference := calculateDifference(float32(humidity), device.MinHumidity, device.MaxHumidity)
		if difference != 0 {
			warninglog := PostgresDBProvider.DB.Model(&Warning{}).Create(WarningLog{
				DeviceSerial: serial,
				DeviceTime:   deviceTime,
				ServerTime:   serverTime,
				Difference:   difference,
			})
			if warninglog.Error != nil {
				return warninglog.Error
			}
		}
	}
	return nil
}

func calculateDifference(humidity float32, minHumidity float32, maxHumidity float32) float32 {
	if maxHumidity < humidity {
		return humidity - maxHumidity
	} else if minHumidity > humidity {
		return minHumidity - humidity
	}
	return 0
}

func CreateDevice(device *Device) error {
	createdData := PostgresDBProvider.DB.Create(&device)
	if createdData.Error != nil {
		log.Info(createdData.Error)
		return createdData.Error
	}
	return nil
}
