package router

import (
	"github.com/andycai/werite/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerCheck = append(routerCheck, registerAdminRouter)
	routerNoCheck = append(routerNoCheck, registerAdminNoCheckRouter)
}

func registerAdminNoCheckRouter(r fiber.Router) {
	// r.Post("/register", handler.User.Register)
	// r.Post("/login", handler.User.Login)
	// r.Post("/exit", handler.User.Exit)
}

func registerAdminRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		admin.Get("/login", handler.Admin.Login)
	}
}
