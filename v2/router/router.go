package router

import (
	"github.com/andycai/werite/v2/middleware"
	"github.com/gofiber/fiber/v2"
)

var (
	routerNoCheck = make([]func(fiber.Router), 0)
	routerCheck   = make([]func(fiber.Router), 0)
)

func Setup(app *fiber.App) {
	v2 := app.Group("/")
	// 需要登录校验的路由
	for _, f := range routerCheck {
		f(v2)
	}

	// v2 := app.Group("/api", JWTAuthMiddleware)
	// 不需要登录校验的路由
	for _, f := range routerNoCheck {
		f(v2)
	}
	app.Use(middleware.NotFoundRoute)
}

// func adminNoCheckRoleRouter(r *gin.Engine) {
// 	// 可根据业务需求来设置接口版本
// 	v1 := r.Group("/api")
// 	// 空接口防止v1定义无使用报错
// 	v1.GET("/nilcheckrole", nil)

// 	for _, f := range routerNoCheckRole {
// 		f(v1)
// 	}
// }

// func adminCheckRoleRouter(r *gin.Engine) {
// 	// 可根据业务需求来设置接口版本
// 	v1 := r.Group("/api", JWTAuthMiddleware)
// 	// 空接口防止v1定义无使用报错
// 	v1.GET("/checkrole", nil)
// 	for _, f := range routerCheckRole {
// 		f(v1)
// 	}
// }

// // JWTAuthMiddleware 基于JWT的认证中间件
// func JWTAuthMiddleware(c *gin.Context) {
// 	//客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
// 	//这里假设Token放在Header的Authorization中，并使用Bearer开头
// 	//这里的具体实现方式要依据你的实际业务情况决定
// 	authHeader := c.Request.Header.Get("Auth")
// 	if authHeader == "" {
// 		net.ResponseError(c, net.CodeLoginExpire)
// 		c.Abort()
// 		return
// 	}

// 	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
// 	token, err := jwt.ParseToken(authHeader)
// 	if err != nil {
// 		logrus.Errorln("token解析失败")
// 		net.ResponseError(c, net.CodeLoginExpire)
// 		c.Abort()
// 		return
// 	}
// 	// 续签 依赖Redis
// 	c.Set(currentUser, token)
// 	c.Set(currentUserName, token.Username)

// 	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
// }
