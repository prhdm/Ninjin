package mqtt

import (
	"encoding/json"
	"farm/src/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Handler struct {
}

type jsonType struct {
	Data []float32 `json:"data"`
}

func NewLogHandler() *Handler {
	return &Handler{}
}

func getTime(deviceTime string) time.Time{
	year, _ := strconv.ParseInt(deviceTime[0:4], 10, 32)
	month, _ := strconv.ParseInt(deviceTime[4:6], 10, 32)
	day, _ := strconv.ParseInt(deviceTime[6:8], 10, 32)
	hour, _ := strconv.ParseInt(deviceTime[8:10], 10, 32)
	minute, _ := strconv.ParseInt(deviceTime[10:12], 10, 32)
	second, _ := strconv.ParseInt(deviceTime[12:14], 10, 32)
	loc, _ := time.LoadLocation("Asia/Tehran")
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
}

func secondTypeHandler(payload string) string{
	sz := len(payload)
	firstPart := payload[sz - 4:]
	if firstPart == "{ee}"{
		payload = payload[:sz - 4]
	}
	return payload
}

func (h *Handler) Handle(payload string, topic string) {
	payload = secondTypeHandler(payload)

	result := jsonType{}
	err := json.Unmarshal([]byte(payload), &result)
	if err != nil{
		log.Info(err)
		return
	}

	serial := topic[9:]
	humidity := int32(result.Data[19])
	deviceTime := getTime(fmt.Sprintf("%d", int64(result.Data[4])))
	serverTime := time.Now()

	models.CreateDeviceLog(serial, humidity, deviceTime, serverTime)
	return
}

