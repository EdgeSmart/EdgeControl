package dev

import (
	"github.com/EdgeSmart/EdgeControl/service/menu"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	ctx.HTML(200, "dev_index.tmpl", gin.H{
		"title": "首页 - 开发者平台",
		"menu":  menu.GetMenu(ctx),
	})
	return

}
