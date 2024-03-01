package device

import (
	"context"
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

func (d DeviceDao) CreateDeviceInfo(ctx context.Context, deviceCode string) (device *model.Device, err error) {
	var de *model.Device
	d.data.Db.Where("device_id=", deviceCode).First(&de)

	if de != nil {
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
