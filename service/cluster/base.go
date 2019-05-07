package cluster

import (
	"encoding/json"

	"github.com/EdgeSmart/EdgeControl/service/manager"
)

type ReqParam struct {
	Token string `json:"token"`
}

type ResParam struct {
	CID   string `json:"cid"`
	CKey  string `json:"ckey"`
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type RetStruct struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []ResParam `json:"data"`
}

func GetList(token string) ([]ResParam, error) {
	retStruct := RetStruct{}
	data, err := manager.Request("POST", "/cluster/list", nil, nil)
	err = json.Unmarshal(data, &retStruct)
	return retStruct.Data, err
}
