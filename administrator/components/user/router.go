package user

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.user.gorm.db"
	KeyNoCheckRouter = "admin.user.router.nocheck"
	KeyCheckRouter   = "admin.user.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		admin.Get("/login", handleLoginPage)
		admin.Post("/login", handleLoginAction)
	}
}

func initCheckRouter(r fiber.Router) {
	r.Get("/logout", handleLogoutAction)
	r.Get("/dashboard", handleDashBoardPage)

	r.Get("/users/profile", handleProfilePage)
	r.Post("/users/profile", handleProfileSave)
	r.Get("/users/security", handleSecurityPage)
	r.Post("/users/password", handlePasswordSave)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterNoCheckRouter(KeyNoCheckRouter, initNoCheckRouter)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
