package mqtt

import (
	"encoding/json"
	"farm/src/models"
	"fmt"
)

type Handler struct {

}

func NewLogHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(payload string) {
	fmt.Println(payload)
	deviceLog := &models.DeviceLog{}
	err := json.Unmarshal([]byte(payload), deviceLog)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(deviceLog)
	models.CreateDataLog(deviceLog)
}

