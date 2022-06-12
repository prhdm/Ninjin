package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func GetLastWarnings(farmId uint) ([]*WarningLog, error) {
	var warningLogSlice []*WarningLog
	deviceSlice, err := GetDeviceList(farmId)

	if err != nil {
		log.Info(err)
		return nil, err
	}

	for _, device := range deviceSlice {
		warningLog := &WarningLog{}
		result := PostgresDBProvider.DB.Where("device_serial = ? AND server_time >= ?", device.DeviceSerial, time.Now().Add(-time.Minute * 5)).Last(warningLog)

		if result.Error != nil {
			if !result.RecordNotFound() {
				return nil, result.Error
			}
			log.Info(result.Error)
		} else {
			warningLogSlice = append(warningLogSlice, warningLog)
		}
	}
	return warningLogSlice, nil
}
