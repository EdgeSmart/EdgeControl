package service

import (
	"encoding/json"
	"net/http"

	"github.com/EdgeSmart/EdgeControl/data/manager"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginReq struct {
	Type     string `json:"type"`
	Identity string `json:"identity"`
	Token    string `json:"token"`
}

type loginRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type managerRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
	MS int64 `json:"ms"`
}

// Login User login
func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.tmpl", gin.H{
		"title": "登录",
	})
	return
}

// LoginCheck User login
func LoginCheck(ctx *gin.Context) {
	req := loginReq{}
	err := ctx.BindJSON(&req)
	if err != nil {
		// todo: add log
	}
	reqData, err := manager.Request("POST", "/user/login", nil, req)
	manager := managerRes{}
	err = json.Unmarshal(reqData, &manager)
	if manager.Data.Token == "" {
		ctx.JSON(http.StatusOK, loginRes{
			Status:  -1,
			Message: "The password is incorrect or the user does not exist.",
		})
		return
	}
	session := sessions.Default(ctx)
	session.Set("token", manager.Data.Token)
	session.Save()
	ctx.JSON(http.StatusOK, loginRes{
		Status:  0,
		Message: "Success",
	})
	return
}

// Logout Logout
func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusFound, "/login")
	return
}
