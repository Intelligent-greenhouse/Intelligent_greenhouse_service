package main

import (
	"github.com/go-kratos/kratos/v2/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"os"

	"intelligent-greenhouse-service/conf"

	"github.com/cypunsource/cypunsource-tool/util"

	_ "github.com/cypunsource/kratos-extension/encoding"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"
)

var (
	Name = "intelligent-greenhouse-service"

	id, _ = os.Hostname()
)

var (
	LogLevel     = log.ParseLevel(util.Getenv("LOG_LEVEL", "DEBUG")) // 日志等级
	AppConfigDir = util.Getenv("APP_CONFIG_DIR", "./data/conf")
)

func main() {
	json.MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true, //默认值不忽略
		UseProtoNames:   true, //使用proto name返回http字段
	}
	conf.InitConfFile(AppConfigDir)

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
	)
	logger = log.With(logger, "level", LogLevel)

	c := config.New(
		config.WithSource(
			file.NewSource(AppConfigDir),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := newApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
