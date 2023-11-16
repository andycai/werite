package main

import (
	"path/filepath"

	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/library/database"
	"github.com/andycai/werite/library/renderer"
	"github.com/andycai/werite/log"
	"github.com/andycai/werite/middlewares"
	"github.com/andycai/werite/v2/dao"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	engine := renderer.ViewEngineStart()
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	log.Setup()
	conf.ReadConf()

	// database open and init
	db, err := database.InitRDBMS(viper.GetString("db.type"), viper.GetString("db.dsn"))
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)
	core.SetupDatabase([]*gorm.DB{db})
	authentication.SessionStart()

	// Middleware
	middlewares.Use(app)

	app.Static("/static", filepath.Join("", "assets"))

	// router
	core.SetupRouter(app)

	err = app.Listen(viper.GetString("httpserver.addr"))
	if err != nil {
		panic(err)
	}
	defer func() {
		//
	}()
}
