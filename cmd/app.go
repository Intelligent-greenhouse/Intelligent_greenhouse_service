package main

import (
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain"
	"intelligent-greenhouse-service/infra"
	"intelligent-greenhouse-service/infra/dao"
	"intelligent-greenhouse-service/service"
	"intelligent-greenhouse-service/trigger"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

func newApp(config *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	dataData, dataCleanup, err := infra.NewData(config.GetSource(), logger)
	if err != nil {
		return nil, nil, err
	}
	repo := dao.NewAuthDao(dataData, logger)
	usecase := domain.NewAuthUsecase(repo, logger)
	service := service.NewAuthService(usecase)
	httpServer := trigger.NewHTTPServer(config.GetTrigger(), service, logger)
	cleanup := func() {
		dataCleanup()
		// 其他的关闭动作
	}

	appInstance := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			httpServer,
		),
	)

	return appInstance, cleanup, nil
}
