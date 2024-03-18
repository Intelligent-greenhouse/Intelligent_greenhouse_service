package device

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/device"
	"intelligent-greenhouse-service/infra"
	"intelligent-greenhouse-service/model"
	"time"
)

type DeviceDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
}

func (d DeviceDao) GetDeviceByDeviceCode(ctx context.Context, deviceCode string) (*model.Device, error) {
	var deviceInfo *model.Device
	err := d.data.Db.Where("device_id = ?", deviceCode).First(&deviceInfo).Error
	if err != nil {
		return nil, err
	}

	return deviceInfo, nil
}

func (d DeviceDao) SetDeviceButton(ctx context.Context, buttonInfo *model.Device) error {
	var deviceInfo *model.Device
	err := d.data.Db.Where("id = ?", buttonInfo.ID).First(&deviceInfo).Error
	if err != nil {
		return err
	}

	if deviceInfo.IsActivation {
		return errors.New(403, "", "deviceMode: auto")
	}

	deviceInfo.Led = buttonInfo.Led
	deviceInfo.Fan = buttonInfo.Fan
	deviceInfo.Water = buttonInfo.Water
	deviceInfo.ChemicalFertilizer = buttonInfo.ChemicalFertilizer
	deviceInfo.IncreaseTemperature = buttonInfo.IncreaseTemperature
	deviceInfo.ReduceTemperature = buttonInfo.ReduceTemperature
	deviceInfo.Buzzer = buttonInfo.Buzzer

	return d.data.Db.Updates(&deviceInfo).Error
}

func (d DeviceDao) SetActiveMode(ctx context.Context, mode bool, deviceId int32) error {
	return d.data.Db.Model(&model.Device{}).Where("id = ?", deviceId).Update("is_activation", mode).Error
}

func (d DeviceDao) UpdateDeviceDes(ctx context.Context, deviceCode, msg string) error {
	fmt.Println("aaa")
	return d.data.Db.Model(&model.Device{}).Where("device_id = ?", deviceCode).Update("des", msg).Error
}

func (d DeviceDao) UpdateDeviceInfo(ctx context.Context, deviceInfo *model.Device) (err error) {
	var od *model.Device
	err = d.data.Db.Where("device_id = ?", deviceInfo.DeviceId).First(&od).Limit(1).Error
	if err != nil {
		return
	}

	deviceInfo.DeviceId = od.DeviceId
	err = d.data.Db.Save(&deviceInfo).Error
	return
}

func (d DeviceDao) GetDeviceList(ctx context.Context, deviceIdList []int32) (deviceList []*model.Device, err error) {
	err = d.data.Db.Where("id in ?", deviceIdList).Find(&deviceList).Error
	return
}

func (d DeviceDao) GetUserDevice(ctx context.Context, deviceId, userId int32) (*model.UserDevice, error) {
	var ud model.UserDevice
	err := d.data.Db.Where("user_id = ? AND device_id = ?", userId, deviceId).First(&ud).Error
	if err != nil {
		return nil, err
	}

	return &ud, nil
}

func (d DeviceDao) CreateDeviceInfo(ctx context.Context, deviceCode string) (device *model.Device, err error) {
	if deviceCode == "" {
		return nil, errors.New(500, "", "name")
	}

	var de *model.Device
	tx := d.data.Db.Where("device_id = ?", deviceCode).First(&de)

	if tx.RowsAffected != 0 {
		return nil, errors.New(409, "", "device code has exist")
	}

	newDevice := &model.Device{
		DeviceId:     deviceCode,
		IsActivation: false,
		RunTime:      time.Now(),
	}
	if err = d.data.Db.Create(&newDevice).Error; err != nil {
		return nil, errors.New(500, "", "db error")
	}

	return newDevice, nil
}

func (d DeviceDao) GetDeviceById(ctx context.Context, deviceId int32) (device *model.Device, err error) {
	var de model.Device
	tx := d.data.Db.First(&de, deviceId)
	err = tx.Error
	if err != nil {
		return nil, errors.New(404, "", "device not found")
	}

	return &de, nil
}

var instance *DeviceDao

func NewDeviceDao(data *infra.Data, c *conf.Bootstrap, logger log.Logger) device.DeviceRepo {
	if instance == nil {
		instance = &DeviceDao{
			data: data,
			log:  log.NewHelper(logger),
			conf: c,
		}
	}
	return instance
}

func GetDeviceDaoInstance() device.DeviceRepo {
	return instance
}
