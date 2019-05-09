package middleware

import (
	"net/http"
	"strings"

	"github.com/EdgeSmart/EdgeControl/service/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginControl login
func LoginControl(ctx *gin.Context) {
	if ctx.Writer.Status() == http.StatusOK {
		if ctx.Request.RequestURI != "/" && !(strings.Index(ctx.Request.RequestURI, "/login") > 0) {
			session := sessions.Default(ctx)
			uidItfc := session.Get("uid")
			uid, ok := uidItfc.(uint64)
			if !ok || uid < 1 {
				if ctx.GetHeader("X-Requested-With") == "XMLHttpRequest" {
					ctx.JSON(http.StatusForbidden, types.ResponseJSON{
						Code:    -1,
						Message: "Need login",
					})
				} else {
					ctx.Redirect(http.StatusFound, "/")
				}
				return
			}
		}
	}
	ctx.Next()
	return
}
