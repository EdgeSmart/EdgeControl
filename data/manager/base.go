package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

// ManageServer ManageServer
type ManageServer struct {
	Host string
	Port uint16
}

var manage = ManageServer{}

func init() {
	getConf()
}

func getConf() {
	appPath, err := os.Getwd()
	if err != nil {
		// todo: add log
	}
	pathSep := string(os.PathSeparator)
	confFile := appPath + pathSep + "conf" + pathSep + "manager.toml"

	_, err = toml.DecodeFile(confFile, &manage)
}

// Request Request server
func Request(method string, url string, query interface{}, post interface{}) ([]byte, error) {
	client := &http.Client{}
	reqURL, _ := getURL(url, query)
	postData, _ := json.Marshal(&post)
	request, err := http.NewRequest(method, reqURL, bytes.NewReader(postData))
	if err != nil {
		// todo: add log
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		// todo: add log
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	return bodyBytes, err
}

func getURL(url string, query interface{}) (string, error) {
	reqURL := ""
	switch query.(type) {
	case nil:
		reqURL = fmt.Sprintf("http://%s:%d%s", manage.Host, manage.Port, url)
	case string:
		reqQuery := query.(string)
		reqURL = fmt.Sprintf("http://%s:%d%s?%s", manage.Host, manage.Port, url, reqQuery)
	}
	return reqURL, nil
}
