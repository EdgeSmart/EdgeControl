package menu

var MainNodes = map[string][]MainNode{
	"/": []MainNode{
		MainNode{
			Text: "个人设置",
			Nodes: []Node{
				Node{
					URL:  "/my",
					Text: "个人信息",
				},
			},
		},
	},
	"user": []MainNode{
		MainNode{
			Text: "集群管理",
			Nodes: []Node{
				Node{
					URL:  "/user/cluster",
					Text: "集群列表",
				},
				Node{
					URL:  "/user/cluster/create",
					Text: "创建集群",
				},
			},
		},
	},
	"dev": []MainNode{
		MainNode{
			Text: "应用管理",
			Nodes: []Node{
				Node{
					URL:  "/dev/app",
					Text: "应用列表",
				},
				Node{
					URL:  "/dev/app/create",
					Text: "创建应用",
				},
			},
		},
	},
	"admin": []MainNode{
		MainNode{
			Text: "用户管理",
			Nodes: []Node{
				Node{
					URL:  "/admin/user",
					Text: "用户列表",
				},
			},
		},
		MainNode{
			Text: "应用管理",
			Nodes: []Node{
				Node{
					URL:  "/admin/app",
					Text: "应用列表",
				},
			},
		},
		MainNode{
			Text: "集群管理",
			Nodes: []Node{
				Node{
					URL:  "/admin/cluster",
					Text: "集群列表",
				},
			},
		},
	},
}
