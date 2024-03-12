package device

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/infra"
	jwt "intelligent-greenhouse-service/middleware"
	"intelligent-greenhouse-service/model"
)

type DeviceRepo interface {
	GetDeviceById(ctx context.Context, deviceId int32) (*model.Device, error)
	GetDeviceByDeviceCode(ctx context.Context, deviceCode string) (*model.Device, error)
	CreateDeviceInfo(ctx context.Context, deviceCode string) (*model.Device, error)
	GetUserDevice(ctx context.Context, deviceId, userId int32) (*model.UserDevice, error)
	GetDeviceList(ctx context.Context, deviceIdList []int32) ([]*model.Device, error)
	UpdateDeviceInfo(ctx context.Context, deviceInfo *model.Device) error
	UpdateDeviceDes(ctx context.Context, deviceCode, msg string) error
	SetActiveMode(ctx context.Context, mode bool, deviceId int32) error
	SetDeviceButton(ctx context.Context, buttonInfo *model.Device) error
}

type DeviceDomain struct {
	mqtt *infra.Mqtt
	repo DeviceRepo
	log  *log.Helper
}

func NewDeviceDomain(repo DeviceRepo, logger log.Logger, m *infra.Mqtt) *DeviceDomain {
	return &DeviceDomain{
		mqtt: m,
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (d DeviceDomain) GetDeviceInfoByDeviceId(ctx context.Context, deviceId int32) (*model.Device, error) {
	return d.repo.GetDeviceById(ctx, deviceId)
}

func (d DeviceDomain) CreateDevice(ctx context.Context, deviceCode string) (*model.Device, error) {
	return d.repo.CreateDeviceInfo(ctx, deviceCode)
}

func (d DeviceDomain) MqttTest(ctx context.Context) {
	d.mqtt.Mq.Publish("service", 0, false, "hello")
}

func (d DeviceDomain) UpdateDeviceInfo(ctx context.Context, deviceInfo *model.Device) error {
	device, err := d.repo.GetDeviceByDeviceCode(ctx, deviceInfo.DeviceId)
	if err != nil {
		return err
	}

	device.Co2 = deviceInfo.Co2
	device.LightIntensity = deviceInfo.LightIntensity
	device.AirTemperature = deviceInfo.AirTemperature
	device.AirHumidity = deviceInfo.AirHumidity
	device.SoilTemperature = deviceInfo.SoilTemperature
	device.SoilMoisture = deviceInfo.SoilMoisture
	device.SoilConductivity = deviceInfo.SoilConductivity
	device.SoilPH = deviceInfo.SoilPH

	device.Led = deviceInfo.Led
	device.Fan = deviceInfo.Fan
	device.Water = deviceInfo.Water
	device.ChemicalFertilizer = deviceInfo.ChemicalFertilizer
	device.IncreaseTemperature = deviceInfo.IncreaseTemperature
	device.ReduceTemperature = deviceInfo.ReduceTemperature
	device.Buzzer = deviceInfo.Buzzer

	if device.IsActivation {
		// CO2>1000ppm
		if deviceInfo.Co2 > 1000 {
			device.Buzzer = true
			device.Fan = true
		}

		//LightIntensity<30000lx
		if deviceInfo.LightIntensity < 30000 {
			device.Led = true
			device.Buzzer = true
		}

		//AirTemperature>28
		if deviceInfo.AirTemperature > 28 {
			device.ReduceTemperature = true
			device.Buzzer = true
		}

		//AirTemperature<20
		if deviceInfo.AirTemperature < 20 {
			device.IncreaseTemperature = true
			device.Buzzer = true
		}

		//SoilMoisture<60%
		if deviceInfo.SoilMoisture < 60 {
			device.Water = true
			device.Buzzer = true
		}

		//CO2>1000ppm && LightIntensity<30000lx
		if deviceInfo.Co2 > 1000 && deviceInfo.LightIntensity < 30000 {
			device.Led = true
			device.Fan = true
			device.Buzzer = true
		}

		//CO2>1000ppm && AirTemperature>28
		if deviceInfo.Co2 > 1000 && deviceInfo.AirTemperature > 28 {
			device.Fan = true
			device.Buzzer = true
		}

		//CO2>1000ppm && AirTemperature<20
		if deviceInfo.Co2 > 1000 && deviceInfo.AirTemperature < 20 {
			device.Fan = true
			device.IncreaseTemperature = true
			device.Buzzer = true
		}

		//CO2>1000ppm && SoilMoisture<60%
		if deviceInfo.Co2 > 1000 && deviceInfo.SoilMoisture < 60 {
			device.Fan = true
			device.Water = true
			device.Buzzer = true
		}

		//LightIntensity<30000lx && AirTemperature>28
		if deviceInfo.LightIntensity < 30000 && deviceInfo.AirTemperature > 28 {
			device.Led = true
			device.ReduceTemperature = true
			device.Buzzer = true
		}

		//LightIntensity<30000lx && AirTemperature<20
		if deviceInfo.LightIntensity < 30000 && device.AirTemperature < 20 {
			device.Led = true
			device.IncreaseTemperature = true
			device.Buzzer = true
		}

		//AirTemperature>28 && SoilMoisture<60%
		if deviceInfo.AirTemperature > 28 && device.SoilMoisture < 60 {
			device.Water = true
			device.ReduceTemperature = true
			device.Buzzer = true
		}

		//AirTemperature<20 && SoilMoisture<60%
		if deviceInfo.AirTemperature < 20 && device.SoilTemperature < 60 {
			device.Water = true
			device.IncreaseTemperature = true
			device.Buzzer = true
		}

		switchInfo, _ := json.Marshal(bodyDeviceButtonInfoPointer(device))
		d.mqtt.Mq.Publish(deviceInfo.DeviceId+"-Command", 0, false, switchInfo)
	}

	return d.repo.UpdateDeviceInfo(ctx, device)
}

func (d DeviceDomain) UpdateDeviceDes(ctx context.Context, deviceCode, msg string) error {
	return d.repo.UpdateDeviceDes(ctx, deviceCode, msg)
}

func (d DeviceDomain) SetDeviceAutoMode(ctx context.Context, mode bool, deviceId int32) error {
	userId, _ := jwt.FromLoginTokenContext(ctx)
	if userId.IsAdmin {
		return d.repo.SetActiveMode(ctx, mode, deviceId)
	}

	_, err := d.repo.GetUserDevice(ctx, deviceId, userId.UserID)
	if err != nil {
		return err
	}

	return d.repo.SetActiveMode(ctx, mode, deviceId)
}

func (d DeviceDomain) SetDeviceButtonSwitch(ctx context.Context, deviceButtonInfo *model.Device) error {
	userId, _ := jwt.FromLoginTokenContext(ctx)
	if userId.IsAdmin {
		return d.repo.SetDeviceButton(ctx, deviceButtonInfo)
	}

	_, err := d.repo.GetUserDevice(ctx, deviceButtonInfo.ID, userId.UserID)
	if err != nil {
		return err
	}

	return d.repo.SetDeviceButton(ctx, deviceButtonInfo)
}

func (d DeviceDomain) GetDeviceMode(ctx context.Context, deviceId int32) (bool, error) {
	id, err := d.repo.GetDeviceById(ctx, deviceId)
	return id.IsActivation, err
}

type DeviceSwitch struct {
	LED                 bool `json:"LED"`
	Fan                 bool `json:"Fan"`
	Water               bool `json:"Water"`
	ChemicalFertilizer  bool `json:"ChemicalFertilizer"`
	IncreaseTemperature bool `json:"IncreaseTemperature"`
	ReduceTemperature   bool `json:"ReduceTemperature"`
	Buzzer              bool `json:"Buzzer"`
}

var bodyDeviceButtonInfoPointer = func(d *model.Device) *DeviceSwitch {
	return &DeviceSwitch{
		LED:                 d.Led,
		Fan:                 d.Fan,
		Water:               d.Water,
		ChemicalFertilizer:  d.ChemicalFertilizer,
		IncreaseTemperature: d.IncreaseTemperature,
		ReduceTemperature:   d.ReduceTemperature,
		Buzzer:              d.Buzzer,
	}
}
