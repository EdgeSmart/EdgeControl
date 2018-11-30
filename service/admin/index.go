package admin

import (
	"github.com/EdgeSmart/EdgeControl/data/menu"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	menu.GetMenu(ctx)
	ctx.HTML(200, "admin_index.tmpl", gin.H{
		"title": "首页 - 管理平台",
		"menu":  menu.GetMenu(ctx),
	})
	return
}
