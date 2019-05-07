package admin

import (
	"fmt"
	"net/http"

	"github.com/EdgeSmart/EdgeControl/service/docker"
	"github.com/gin-gonic/gin"
)

// ClusterList ClusterList
func ClusterList(ctx *gin.Context) {
	clientCtx, dockerClient, err := docker.New("cluster_test")
	if err != nil {
		fmt.Println(err)
		return
	}
	aaa, err := dockerClient.Info(clientCtx)
	fmt.Println(aaa, err)
	ctx.JSON(http.StatusOK, aaa)
}
