package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdgeSmart/EdgeControl/service/cluster"
	"github.com/EdgeSmart/EdgeControl/service/menu"
	"github.com/gin-gonic/gin"
)

func ClusterList(ctx *gin.Context) {

	data, _ := cluster.GetList("")
	dataByte, _ := json.Marshal(data)
	fmt.Println(string(dataByte))
	ctx.HTML(http.StatusOK, "user_cluster.tmpl", gin.H{
		"title": "集群列表 - 用户平台",
		"menu":  menu.GetMenu(ctx),
		"data":  string(dataByte),
	})
}

func ClusterView(ctx *gin.Context) {
	cid, exists := ctx.Params.Get("cid")
	if exists {
		fmt.Println(cid)
	}
}
