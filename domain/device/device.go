package device

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/infra"
	"intelligent-greenhouse-service/model"
)

type DeviceRepo interface {
	GetDeviceById(ctx context.Context, deviceId int32) (*model.Device, error)
	CreateDeviceInfo(ctx context.Context, deviceCode string) (*model.Device, error)
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
