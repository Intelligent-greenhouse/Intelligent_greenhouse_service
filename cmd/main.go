package main

import (
	"os"

	"intelligent-greenhouse-service/conf"

	"github.com/cypunsource/cypunsource-tool/util"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"

	_ "github.com/cypunsource/kratos-extension/encoding"
	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "cypunsource-auth"
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

var (
	LogLevel     = log.ParseLevel(util.Getenv("LOG_LEVEL", "DEBUG")) // 日志等级
	AppConfigDir = util.Getenv("APP_CONFIG_DIR", "./data/conf")
)

func main() {
	conf.InitConfFile(AppConfigDir)

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"service.name", Name,
		"service.version", Version,

		"caller", log.DefaultCaller,
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
