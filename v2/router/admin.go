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
	// 注册
	// r.Post("/register", handler.User.Register)
	// 登录
	// r.Post("/login", handler.User.Login)
	// 退出
	// r.Post("/exit", handler.User.Exit)
}

func registerAdminRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		// 获取群组详细信息
		admin.Get("/login", handler.Admin.Login)

	}
}
