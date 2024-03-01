package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	api "intelligent-greenhouse-service/api/web/user"
	"intelligent-greenhouse-service/conf"
	domain "intelligent-greenhouse-service/domain/user"
	"intelligent-greenhouse-service/infra"
	deviceDao "intelligent-greenhouse-service/infra/dao/device"
	"intelligent-greenhouse-service/infra/dao/user"
	userDao "intelligent-greenhouse-service/infra/dao/user"
	service "intelligent-greenhouse-service/service/user"
	"intelligent-greenhouse-service/trigger"
)

func newApp(config *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	data, dataCleanup, err := infra.NewData(config.GetSource(), logger)
	if err != nil {
		return nil, nil, err
	}

	userDao.NewUserDao(data, config, logger)
	deviceDao.NewDeviceDao(data, config, logger)

	httpServer := trigger.NewHTTPServer(config.GetTrigger(), config.Jwt, logger)
	cleanup := func() {
		dataCleanup()
		// 其他的关闭动作
	}

	usecase := domain.NewUserDomain(user.GetUserDaoInstance(), deviceDao.GetDeviceDaoInstance(), logger)
	srv := service.NewUserService(usecase)
	api.RegisterUserHTTPServer(httpServer, srv)

	appInstance := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			httpServer,
		),
	)

	return appInstance, cleanup, nil
}
