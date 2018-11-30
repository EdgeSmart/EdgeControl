package menu

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Node struct {
	URL  string
	Text string
}

type MainNode struct {
	Text  string
	Nodes []Node
}

func GetMenu(ctx *gin.Context) []MainNode {
	menu := []MainNode{}
	path := ctx.Request.URL.Path
	pathArr := strings.Split(path, "/")
	switch pathArr[1] {
	case "user":
		menu = MainNodes["user"]
	case "dev":
		menu = MainNodes["dev"]
	case "admin":
		menu = MainNodes["admin"]
	default:
		menu = MainNodes["/"]
	}
	return menu
}
