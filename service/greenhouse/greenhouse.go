package greenhouse

import (
	"context"
	greenhouseapi "intelligent-greenhouse-service/api/web/greenhouse"
	domain_greenhouse "intelligent-greenhouse-service/domain/greenhouse"
	"intelligent-greenhouse-service/model"
)

type GreenhouseService struct {
	uc *domain_greenhouse.GreenhouseDomain
}

func NewUserService(uc *domain_greenhouse.GreenhouseDomain) *GreenhouseService {
	return &GreenhouseService{uc: uc}
}

func (g GreenhouseService) CreateGreenhouse(ctx context.Context, info *greenhouseapi.CreateGreenhouseInfo) (*greenhouseapi.CreateGreenhouseReply, error) {
	greenhouse, err := g.uc.CreateGreenhouse(ctx, info.Size, info.Name, info.Pos)
	if err != nil {
		return nil, err
	}

	return &greenhouseapi.CreateGreenhouseReply{GreenhouseId: greenhouse.ID}, nil
}

func (g GreenhouseService) GetGreenhouseInfoByUserId(ctx context.Context, request *greenhouseapi.GetGreenhouseInfoByUserIdRequest) (*greenhouseapi.GetGreenhouseInfoByUserIdReply, error) {
	list, err := g.uc.GetGreenhouseListByUserId(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	var greenHouseList []*greenhouseapi.GreenhouseInfo
	for _, g := range list {
		greenHouseList = append(greenHouseList, &greenhouseapi.GreenhouseInfo{
			GreenhouseId: g.ID,
			Pos:          g.Pos,
			Size:         g.Size,
		})
	}

	return &greenhouseapi.GetGreenhouseInfoByUserIdReply{List: greenHouseList}, nil
}

func (g GreenhouseService) GetGreenhouseAllDeviceInfo(ctx context.Context, id *greenhouseapi.GreenhouseId) (*greenhouseapi.GetGreenhouseAllDeviceInfoReply, error) {
	deviceInfo, err := g.uc.GetDeviceListByGreenhouseId(ctx, id.GreenhouseId)
	if err != nil {
		return nil, err
	}

	var deviceList []*greenhouseapi.DeviceInfo
	for _, d := range deviceInfo {
		deviceList = append(deviceList, deviceInfoPointer(d))
	}

	return &greenhouseapi.GetGreenhouseAllDeviceInfoReply{List: deviceList}, nil
}

func (g GreenhouseService) GetDeviceInfo(ctx context.Context, id *greenhouseapi.DeviceId) (*greenhouseapi.DeviceInfo, error) {
	deviceInfo, err := g.uc.GetDeviceInfoByDeviceId(ctx, id.DeviceId)
	if err != nil {
		return nil, err
	}

	return deviceInfoPointer(deviceInfo), nil
}

var deviceInfoPointer = func(d *model.Device) *greenhouseapi.DeviceInfo {
	return &greenhouseapi.DeviceInfo{
		DeviceCode:          d.DeviceId,
		Co2:                 d.SoilConductivity,
		LightIntensity:      d.LightIntensity,
		AirTemperature:      d.AirTemperature,
		AirHumidity:         d.AirHumidity,
		SoilTemperature:     d.SoilPH,
		SoilMoisture:        d.SoilMoisture,
		SoilConductivity:    d.SoilConductivity,
		SoilPh:              d.SoilPH,
		Led:                 d.Led,
		Fan:                 d.Fan,
		Water:               d.Water,
		ChemicalFertilizer:  d.ChemicalFertilizer,
		IncreaseTemperature: d.IncreaseTemperature,
		ReduceTemperature:   d.ReduceTemperature,
		Buzzer:              d.Buzzer,
		Des:                 d.Des,
	}
}
