package index

import (
	"github.com/EdgeSmart/EdgeControl/service/menu"
	"github.com/gin-gonic/gin"
)

// Index Index
func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"title": "首页",
		"menu":  menu.GetMenu(ctx),
	})
	return

}
