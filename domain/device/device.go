package device

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/infra"
	jwt "intelligent-greenhouse-service/middleware"
	"intelligent-greenhouse-service/model"
)

type DeviceRepo interface {
	GetDeviceById(ctx context.Context, deviceId int32) (*model.Device, error)
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

	// TODO 智能/手动更新设备状态

	return d.repo.UpdateDeviceInfo(ctx, deviceInfo)
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
