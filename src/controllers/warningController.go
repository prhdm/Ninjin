package controllers

import (
	"context"
	"farm/src/models"
	pb_warning "farm/src/proto/messages/warning"
)

func (f FarmServer) GetWarnings(ctx context.Context, request *pb_warning.WarningRequest) (*pb_warning.WarningResponse, error) {
	result, err := models.GetLastWarnings(ctx.Value("farm_id").(uint))
	if err != nil {
		return nil, err
	}
	var warningLogSlice []*pb_warning.Warning
	for _, warningLog := range result {
		warningLogPb := pb_warning.Warning{
			DeviceSerial: warningLog.DeviceSerial,
			Difference: warningLog.Difference,
		}
		warningLogSlice = append(warningLogSlice, &warningLogPb)
	}
	return &pb_warning.WarningResponse{
		Warnings: warningLogSlice,
	}, nil
}

