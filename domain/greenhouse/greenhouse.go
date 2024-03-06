package greenhouse

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/domain/device"
	jwt "intelligent-greenhouse-service/middleware"
	"intelligent-greenhouse-service/model"
)

type GreenhouseRepo interface {
	GetGreenhouseListByUserId(ctx context.Context, userId int32) ([]*model.Greenhouse, error)
	GetGreenhouseInfoById(ctx context.Context, id int32) (*model.Greenhouse, error)
	GetGreenhouseBandInfo(ctx context.Context, deviceId, greenhouseId int32) error
	BandGreenhouseAndDevice(ctx context.Context, deviceId, greenhouseId, userId int32) error
	CreateGreenhouse(ctx context.Context, size int32, name, pos string) (*model.Greenhouse, error)
}

type GreenhouseDomain struct {
	greenhouseRepo GreenhouseRepo
	deviceRepo     device.DeviceRepo
	log            *log.Helper
}

func NewGreenhouseDomain(repo GreenhouseRepo, logger log.Logger, de device.DeviceRepo) *GreenhouseDomain {
	return &GreenhouseDomain{
		greenhouseRepo: repo,
		deviceRepo:     de,
		log:            log.NewHelper(logger),
	}
}

func (g GreenhouseDomain) CreateGreenhouse(ctx context.Context, size int32, name, pos string) (*model.Greenhouse, error) {
	return g.greenhouseRepo.CreateGreenhouse(ctx, size, name, pos)
}

func (g GreenhouseDomain) GetGreenhouseListByUserId(ctx context.Context, userId int32) ([]*model.Greenhouse, error) {
	return g.greenhouseRepo.GetGreenhouseListByUserId(ctx, userId)
}

func (g GreenhouseDomain) GetDeviceInfoByDeviceId(ctx context.Context, deviceId int32) (*model.Device, error) {
	userInfo, _ := jwt.FromLoginTokenContext(ctx)
	if userInfo.IsAdmin {
		return g.deviceRepo.GetDeviceById(ctx, deviceId)
	}

	_, err := g.deviceRepo.GetUserDevice(ctx, deviceId, userInfo.UserID)
	if err != nil {
		return nil, err
	}

	return g.deviceRepo.GetDeviceById(ctx, deviceId)
}
