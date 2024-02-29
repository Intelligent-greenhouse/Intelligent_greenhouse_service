package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/infra"
	"intelligent-greenhouse-service/trigger"
)

func Register(
	http *http.Server,
	config *conf.Bootstrap,
	data *infra.Data,
	logger log.Logger,
	handles ...RegisterDomainHandler,
) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ConfigKey{}, config)
	ctx = context.WithValue(ctx, DataKey{}, data)
	ctx = context.WithValue(ctx, LoggerKey{}, logger)
	for _, f := range handles {
		f(ctx, http)
	}
}

func newApp(config *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	data, dataCleanup, err := infra.NewData(config.GetSource(), logger)
	if err != nil {
		return nil, nil, err
	}

	httpServer := trigger.NewHTTPServer(config.GetTrigger(), config.Jwt, logger)
	cleanup := func() {
		dataCleanup()
		// 其他的关闭动作
	}

	Register(
		httpServer,
		config,
		data,
		logger,
		UserRegister(),
	)

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

type RegisterDomainHandler func(ctx context.Context, s *http.Server)

type RegisterDomain func() RegisterDomainHandler

type LoggerKey struct{}

func Logger(ctx context.Context) log.Logger {
	v, ok := ctx.Value(LoggerKey{}).(log.Logger)
	if !ok {
		panic("")
	}
	return v
}

type ConfigKey struct{}

func Config(ctx context.Context) *conf.Bootstrap {
	v, ok := ctx.Value(ConfigKey{}).(*conf.Bootstrap)
	if !ok {
		panic("")
	}
	return v
}

type DataKey struct{}

func Data(ctx context.Context) *infra.Data {
	v, ok := ctx.Value(DataKey{}).(*infra.Data)
	if !ok {
		panic("")
	}
	return v
}
