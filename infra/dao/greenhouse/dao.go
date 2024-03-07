package greenhouse

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/greenhouse"
	"intelligent-greenhouse-service/infra"
	jwt "intelligent-greenhouse-service/middleware"
	"intelligent-greenhouse-service/model"
)

type GreenhouseDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
}

func (g GreenhouseDao) GetDeviceIdListByGreenhouseId(ctx context.Context, greenhouseId int32) ([]int32, error) {
	var idList []int32
	err := g.data.Db.Model(&model.GreenhouseDevice{}).Where("greenhouse_id = ?", greenhouseId).Pluck("device_id", &idList).Error
	if err != nil {
		return nil, err
	}

	return idList, nil
}

func (g GreenhouseDao) CreateGreenhouse(ctx context.Context, size int32, name, pos string) (*model.Greenhouse, error) {
	newGreenhouse := &model.Greenhouse{
		Name: name,
		Pos:  pos,
		Size: size,
	}

	err := g.data.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&newGreenhouse).Error
		if err != nil {
			return err
		}

		userInfo, _ := jwt.FromLoginTokenContext(ctx)
		err = tx.Create(&model.UserGreenhouse{
			GreenhouseId: newGreenhouse.ID,
			UserId:       userInfo.UserID,
		}).Error

		return err
	})
	if err != nil {
		return nil, err
	}

	return newGreenhouse, nil
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

func (g GreenhouseDao) BandGreenhouseAndDevice(ctx context.Context, deviceId, greenhouseId, userId int32) error {
	err := g.data.Db.Transaction(func(tx *gorm.DB) (err error) {

		// 查找用户是否拥对应的大棚数据
		err = tx.Where(&model.UserGreenhouse{GreenhouseId: greenhouseId, UserId: userId}).First(&model.UserGreenhouse{}).Limit(1).Error
		if err != nil {
			return err
		}

		// 查找设备是否已经被绑定
		var existingDevice model.GreenhouseDevice
		if err = tx.Where(&model.GreenhouseDevice{DeviceId: deviceId}).Limit(1).First(&existingDevice).Error; err == nil {
			// 设备已绑定
			return errors.New(409, "", "Device already bound")
		}

		var existingUserDevice model.UserDevice
		if err = tx.Where(&model.UserDevice{DeviceId: deviceId}).Limit(1).First(&existingUserDevice).Error; err == nil {
			// 用户和设备的绑定关系已存在
			return errors.New(409, "", "User and device binding already exists")
		}

		// 创建用户和设备的绑定关系
		if err = tx.Create(&model.UserDevice{UserId: userId, DeviceId: deviceId}).Error; err != nil {
			// 创建失败，回滚事务
			return err
		}

		// 创建设备和温室的绑定关系
		if err = tx.Create(&model.GreenhouseDevice{GreenhouseId: greenhouseId, DeviceId: deviceId}).Error; err != nil {
			// 创建失败，回滚事务
			return err
		}

		if err = tx.Model(&model.Device{}).Where("id = ?", deviceId).Update("is_activation", true).Error; err != nil {
			// 更新失败，回滚事务
			return err
		}

		// 提交事务
		return nil
	})

	return err
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
