package user

import (
	"errors"
	"net/http"

	"github.com/EdgeSmart/EdgeControl/dao"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginInfo struct {
	Type     string `json:"type"`
	Identity string `json:"identity"`
	Token    string `json:"token"`
}

type loginResp struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
	var (
		httpStatus = http.StatusForbidden
		httResp    = loginResp{
			Status:  -1,
			Message: "Login failed",
		}
	)
	defer func() {
		ctx.JSON(httpStatus, httResp)
	}()
	loginData := loginInfo{}
	err := ctx.BindJSON(&loginData)
	if err != nil {
		return
	}
	// login
	uid, err := DoLogin(loginData)
	if err != nil {
		return
	}
	session := sessions.Default(ctx)
	session.Set("uid", uid)
	session.Save()
	httpStatus = http.StatusOK
	httResp = loginResp{
		Status:  0,
		Message: "Success",
		Data: map[string]string{
			"name":   "peter",
			"avatar": "https://raw.githubusercontent.com/taylorchen709/markdown-images/master/vueadmin/user.png",
		},
	}
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

// DoLogin DoLogin
func DoLogin(data loginInfo) (uint64, error) {
	var (
		err error
		uid uint64
	)
	// login type
	switch data.Type {
	case "test":
		uid, err = loginTest(data)
	}
	return uid, err
}

// loginTest loginTest
func loginTest(data loginInfo) (uint64, error) {
	db, _ := dao.GetDB("edge")
	stmt, _ := db.Prepare("SELECT `uid`,`token`,`ext` FROM `user_auth` WHERE `identity` = ? AND `type` = ? AND `status` = ?")
	defer stmt.Close()

	rows := stmt.QueryRow(data.Identity, data.Type, 0)

	var uid uint64
	var token string
	var ext string

	err := rows.Scan(&uid, &token, &ext)
	if err != nil {
		return 0, err
	}
	if token != data.Token {
		return 0, errors.New("Verification failed")
	}
	return uid, nil
}
