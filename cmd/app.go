package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	deviceAPI "intelligent-greenhouse-service/api/mq/device"
	greenhouseAPI "intelligent-greenhouse-service/api/web/greenhouse"
	userAPI "intelligent-greenhouse-service/api/web/user"
	"intelligent-greenhouse-service/conf"
	deviceDomain "intelligent-greenhouse-service/domain/device"
	greenhouseDomain "intelligent-greenhouse-service/domain/greenhouse"
	userDomain "intelligent-greenhouse-service/domain/user"
	"intelligent-greenhouse-service/infra"
	deviceDao "intelligent-greenhouse-service/infra/dao/device"
	greenhouseDao "intelligent-greenhouse-service/infra/dao/greenhouse"
	"intelligent-greenhouse-service/infra/dao/user"
	userDao "intelligent-greenhouse-service/infra/dao/user"

	deviceService "intelligent-greenhouse-service/service/device"
	greenService "intelligent-greenhouse-service/service/greenhouse"
	userService "intelligent-greenhouse-service/service/user"
	"intelligent-greenhouse-service/trigger"
)

func newApp(config *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	data, dataCleanup, err := infra.NewData(config.GetSource(), logger)
	if err != nil {
		return nil, nil, err
	}

	userDao.NewUserDao(data, config, logger)
	deviceDao.NewDeviceDao(data, config, logger)
	greenhouseDao.NewGreenhouseDao(data, config, logger)

	httpServer := trigger.NewHTTPServer(config.GetTrigger(), config.Jwt, logger)
	cleanup := func() {
		dataCleanup()
		// 其他的关闭动作
	}

	userCase := userDomain.NewUserDomain(user.GetUserDaoInstance(), deviceDao.GetDeviceDaoInstance(), greenhouseDao.GetGreenhouseDaoInstance(), logger)
	s1 := userService.NewUserService(userCase)
	userAPI.RegisterUserHTTPServer(httpServer, s1)

	greenhouseCase := greenhouseDomain.NewGreenhouseDomain(greenhouseDao.GetGreenhouseDaoInstance(), logger)
	s2 := greenService.NewUserService(greenhouseCase)
	greenhouseAPI.RegisterGreenhouseHTTPServer(httpServer, s2)

	deviceCase := deviceDomain.NewDeviceDomain(deviceDao.GetDeviceDaoInstance(), logger)
	s3 := deviceService.NewDeviceService(deviceCase)
	deviceAPI.RegisterDeviceHTTPServer(httpServer, s3)

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
