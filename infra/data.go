package infra

import (
	"fmt"
	"gorm.io/driver/mysql"
	"strings"

	"intelligent-greenhouse-service/conf"

	"github.com/cypunsource/cypunsource-tool/database"
	grom_logger "github.com/cypunsource/cypunsource-tool/logger/grom"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// Data .
type Data struct {
	Db *gorm.DB
}

// NewData .
func NewData(c *conf.Source, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db, err := NewDataBase(c.GetDatabase(), logger)
	if err != nil {
		return nil, nil, err
	}
	return &Data{Db: db}, cleanup, nil
}

// NewDataBase 建立数据库连接
func NewDataBase(c *conf.Database, logger log.Logger) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), c.GetDbname())
	dsn = dsn + strings.Join(c.GetOption(), "&")

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: grom_logger.NewLoggerHelper(logger),
	})
	if err != nil {
		return
	}

	selDb, err := db.DB()
	if err != nil {
		return
	}
	// 设置连接池
	selDb.SetMaxIdleConns(int(c.MaxIdleConn))
	selDb.SetMaxOpenConns(int(c.MaxOpenConn))
	// 设置超时 time.Second * 30
	selDb.SetConnMaxLifetime(c.ConnMaxLifetime.AsDuration())

	// 注册钩子
	database.RegisterCallbacks(db)
	return
}
