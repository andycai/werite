package main

import (
	"path/filepath"

	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/library/database/gorm"
	"github.com/andycai/werite/log"
	"github.com/andycai/werite/v2/dao"
	"github.com/andycai/werite/v2/middleware"
	"github.com/andycai/werite/v2/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	log.Setup()
	conf.ReadConf() // 读取配置
	dbType := viper.GetString("db.type")
	dbDSN := viper.GetString("db.dsn")
	db, err := gorm.InitRDBMS(dbType, dbDSN)
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)
	middleware.Use(app) // 初始化中间件

	app.Static("/static", filepath.Join("", "assets"))

	router.Setup(app) // 初始化路由

	err = app.Listen(viper.GetString("httpserver.addr"))
	if err != nil {
		panic(err)
	}
	defer func() {
		//
	}()
}
