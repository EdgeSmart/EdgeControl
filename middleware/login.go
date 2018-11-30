package middleware

import (
	"net/http"
	"strings"

	"github.com/EdgeSmart/EdgeControl/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginControl 登陆控制
func LoginControl(ctx *gin.Context) {
	if ctx.Writer.Status() == 200 {
		if strings.Index(ctx.Request.RequestURI, "/login") != 0 {
			session := sessions.Default(ctx)
			token := session.Get("token")
			switch token.(type) {
			case string:
				tokenStr := token.(string)
				if tokenStr != "" {
					ctx.Next()
					return
				}
			}
			service.Login(ctx)
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
	}
	ctx.Next()
	return
}