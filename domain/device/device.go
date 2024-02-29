package device

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type DeviceRepo interface {
}

type DeviceDomain struct {
	repo DeviceRepo
	log  *log.Helper
}

func NewDeviceDomain(repo DeviceRepo, logger log.Logger) *DeviceDomain {
	return &DeviceDomain{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (d DeviceDomain) GetDeviceInfoByDeviceCode(ctx context.Context) {

}
