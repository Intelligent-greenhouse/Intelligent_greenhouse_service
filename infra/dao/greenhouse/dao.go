package greenhouse

import (
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/greenhouse"
	"intelligent-greenhouse-service/infra"
)

type GreenhouseDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
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
