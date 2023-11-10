package main

import (
	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/library/database/gorm"
	"github.com/andycai/werite/log"
	"github.com/andycai/werite/v2/dao"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	log.Setup()
	conf.ReadConf() // 读取配置
	dbType := viper.GetString("db.type")
	dbDSN := viper.GetString("db.dsn")
	db, err := gorm.InitRDBMS(dbType, dbDSN)
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)
	// middleware.Use(app) // 初始化中间件
	// router.Setup(app)   // 初始化路由
	// system.Cache.InitIds() // 初始化数据表的 ID 到缓存

	err = app.Listen(viper.GetString("httpserver.addr"))
	if err != nil {
		panic(err)
	}
	defer func() {
		//
	}()
}
