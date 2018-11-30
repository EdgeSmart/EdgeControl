package main

import (
	"github.com/EdgeSmart/EdgeControl/middleware"
	"github.com/EdgeSmart/EdgeControl/service"
	"github.com/EdgeSmart/EdgeControl/service/admin"
	"github.com/EdgeSmart/EdgeControl/service/dev"
	"github.com/EdgeSmart/EdgeControl/service/index"
	"github.com/EdgeSmart/EdgeControl/service/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Static("/statics", "./templates/statics")
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("session_id", store))
	app.Use(middleware.LoginControl)
	// index
	app.Any("/", index.Index)
	// login
	app.GET("/login", service.Login)
	app.POST("/login", service.LoginCheck)
	app.GET("/logout", service.Logout)

	// admin
	adminGroup := app.Group("/admin")
	adminGroup.GET("/", admin.Index)

	// user
	userGroup := app.Group("/user")
	userGroup.GET("/", user.Index)
	userGroup.GET("/cluster", user.ClusterList)
	userGroup.GET("/cluster/view/:cid", user.ClusterView)

	// dev
	devGroup := app.Group("/dev")
	devGroup.GET("/", dev.Index)

	app.LoadHTMLGlob("./templates/*/*.tmpl")

	app.Run(":8000")
}
