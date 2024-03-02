package greenhouse

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/greenhouse"
	"intelligent-greenhouse-service/infra"
	"intelligent-greenhouse-service/model"
)

type GreenhouseDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
}

func (g GreenhouseDao) GetGreenhouseListByUserId(ctx context.Context, userId int32) ([]*model.Greenhouse, error) {
	var greenList []*model.Greenhouse
	err := g.data.Db.Table("user_greenhouse").
		Select("greenhouse.*").
		Joins("JOIN greenhouse ON user_greenhouse.greenhouse_id = greenhouse.id").
		Where("user_greenhouse.user_id = ?", userId).
		Find(&greenList).
		Error
	if err != nil {
		return nil, err
	}
	return greenList, nil
}

func (g GreenhouseDao) BandGreenhouseAndDevice(ctx context.Context, deviceId, greenhouseId int32) error {
	err := g.data.Db.Find(&model.GreenhouseDevice{GreenhouseId: greenhouseId, DeviceId: deviceId}).Error
	if err != nil {
		return nil
	}

	return g.data.Db.Create(&model.GreenhouseDevice{GreenhouseId: greenhouseId, DeviceId: deviceId}).Error
}

func (g GreenhouseDao) GetGreenhouseBandInfo(ctx context.Context, deviceId, greenhouseId int32) error {
	return g.data.Db.Find(&model.GreenhouseDevice{GreenhouseId: greenhouseId, DeviceId: deviceId}).Error
}

func (g GreenhouseDao) GetGreenhouseInfoById(ctx context.Context, id int32) (info *model.Greenhouse, err error) {
	err = g.data.Db.Find(&info, id).Error
	return
}

var instance *GreenhouseDao

func NewGreenhouseDao(data *infra.Data, c *conf.Bootstrap, logger log.Logger) greenhouse.GreenhouseRepo {
	if instance == nil {
		instance = &GreenhouseDao{
			data: data,
			log:  log.NewHelper(logger),
			conf: c,
		}
	}
	return instance
}

func GetGreenhouseDaoInstance() greenhouse.GreenhouseRepo {
	return instance
}
