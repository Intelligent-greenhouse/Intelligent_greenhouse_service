package greenhouse

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/model"
)

type GreenhouseRepo interface {
	GetGreenhouseListByUserId(ctx context.Context, userId int32) ([]*model.Greenhouse, error)
	GetGreenhouseInfoById(ctx context.Context, id int32) (*model.Greenhouse, error)
	GetGreenhouseBandInfo(ctx context.Context, deviceId, greenhouseId int32) error
	BandGreenhouseAndDevice(ctx context.Context, deviceId, greenhouseId int32) error
	CreateGreenhouse(ctx context.Context, size int32, name, pos string) (*model.Greenhouse, error)
}

type GreenhouseDomain struct {
	greenhouseRepo GreenhouseRepo
	log            *log.Helper
}

func NewGreenhouseDomain(repo GreenhouseRepo, logger log.Logger) *GreenhouseDomain {
	return &GreenhouseDomain{
		greenhouseRepo: repo,
		log:            log.NewHelper(logger),
	}
}

func (g GreenhouseDomain) CreateGreenhouse(ctx context.Context, size int32, name, pos string) (*model.Greenhouse, error) {
	return g.greenhouseRepo.CreateGreenhouse(ctx, size, name, pos)
}

func (g GreenhouseDomain) GetGreenhouseListByUserId(ctx context.Context, userId int32) ([]*model.Greenhouse, error) {
	return g.greenhouseRepo.GetGreenhouseListByUserId(ctx, userId)
}
