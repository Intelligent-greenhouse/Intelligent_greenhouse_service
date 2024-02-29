package device

import (
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/device"
	"intelligent-greenhouse-service/infra"
)

type DeviceDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
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
