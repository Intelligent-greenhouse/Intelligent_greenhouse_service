package device

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "intelligent-greenhouse-service/api/mq/device"
	domain_device "intelligent-greenhouse-service/domain/device"
	"intelligent-greenhouse-service/model"
	"time"
)

type DeviceService struct {
	uc *domain_device.DeviceDomain
}

func NewDeviceService(uc *domain_device.DeviceDomain) *DeviceService {
	return &DeviceService{uc: uc}
}

func (d DeviceService) UpdateDeviceInfo(ctx context.Context, request *v1.UpdateDeviceInfoRequest) (*emptypb.Empty, error) {
	err := d.uc.UpdateDeviceInfo(ctx, bodyDevicePointer(request))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (d DeviceService) UpdateDeviceDes(ctx context.Context, request *v1.DeviceDes) (*emptypb.Empty, error) {
	err := d.uc.UpdateDeviceDes(ctx, request.DeviceCode, request.Msg)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

var bodyDevicePointer = func(d *v1.UpdateDeviceInfoRequest) *model.Device {
	return &model.Device{
		DeviceId:            d.DeviceCode,
		Co2:                 d.Co2,
		LightIntensity:      d.LightIntensity,
		AirTemperature:      d.AirTemperature,
		AirHumidity:         d.AirHumidity,
		SoilTemperature:     d.SoilTemperature,
		SoilMoisture:        d.SoilMoisture,
		SoilConductivity:    d.SoilConductivity,
		SoilPH:              d.SoilPh,
		Led:                 d.Led,
		Fan:                 d.Fan,
		Water:               d.Water,
		ChemicalFertilizer:  d.ChemicalFertilizer,
		IncreaseTemperature: d.IncreaseTemperature,
		ReduceTemperature:   d.ReduceTemperature,
		Buzzer:              d.Buzzer,
		RunTime:             time.Time{},
	}
}
