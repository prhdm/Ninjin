package mqtt

import (
	"encoding/json"
	"fmt"
)

type Handler struct {
}

type jsonType struct {
	Data []float32 `json:"data"`
}

func NewLogHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(payload string) {
	fmt.Println(payload)
	result := jsonType{}
	err := json.Unmarshal([]byte(payload), &result)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	//models.CreateDataLog(deviceLog)
}

