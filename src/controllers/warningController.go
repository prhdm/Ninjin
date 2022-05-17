package controllers

import (
	pb_warning "farm/src/proto/messages/warning"
	"farm/src/proto/service/farm"
)

func (f FarmServer) Warning(request *pb_warning.WarningRequest, stream farm.Farm_WarningServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			err := stream.SendMsg(&pb_warning.WarningResponse{Difference: float32(-1)})
			if err != nil {
				return err
			}
		}
	}
}
