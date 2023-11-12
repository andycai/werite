package main

import (
	"path/filepath"

	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/library/database/gorm"
	"github.com/andycai/werite/library/renderer"
	"github.com/andycai/werite/log"
	"github.com/andycai/werite/v2/dao"
	"github.com/andycai/werite/v2/middleware"
	"github.com/andycai/werite/v2/router"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: renderer.ViewEngineStart(),
	})

	log.Setup()
	conf.ReadConf()

	// database open and init
	db, err := gorm.InitRDBMS(viper.GetString("db.type"), viper.GetString("db.dsn"))
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)
	authentication.SessionStart()

	// Middleware
	middleware.Use(app)

	app.Static("/static", filepath.Join("", "assets"))

	// router
	router.Setup(app)

	err = app.Listen(viper.GetString("httpserver.addr"))
	if err != nil {
		panic(err)
	}
	defer func() {
		//
	}()
}
