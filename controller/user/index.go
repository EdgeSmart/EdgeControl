package user

import (
	"github.com/EdgeSmart/EdgeControl/service/menu"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.HTML(200, "user_index.tmpl", gin.H{
		"title": "首页 - 用户平台",
		"menu":  menu.GetMenu(ctx),
	})
	return

}
