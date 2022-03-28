package main

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/model/response"
	"io/ioutil"
	"net/http"
	"strings"
)

//type getHostModel struct {
//	Jsonrpc string `json:"jsonrpc"`
//	Method  string `json:"method"`
//	Params  struct {
//		Output           []string `json:"output"`
//		SelectInterfaces []string `json:"selectInterfaces"`
//	} `json:"params"`
//	ID   int    `json:"id"`
//	Auth string `json:"auth"`
//}

type ResultModel struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

func PostProcessingJSON(ProcessStruct interface{}, url string) (Jsonstr string) {
	jsonBytes, _ := json.Marshal(ProcessStruct)
	body := string(jsonBytes)
	response, err := http.Post(url, "application/json-rpc; charset=utf-8", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	JsonStr := string(content)
	return JsonStr
}

func ZabbixLogin(username string, password string, url string) (resultToken string) {
	var resultModel ResultModel
	var loginstrcut = response.ZabbixLoginModel{
		Jsonrpc: "2.0",
		Method:  "user.login",
		Params: struct {
			User     string `json:"user"`
			Password string `json:"password"`
		}(struct {
			User     string
			Password string
		}{User: username, Password: password}),
	}
	json.Unmarshal([]byte(PostProcessingJSON(loginstrcut, url)), &resultModel)
	resultToken = resultModel.Result
	return resultToken
}

func GetHostList() (r1 string) {
	//var resultModel ResultModel
	var outPutList []string
	var interfaceList []string
	outPutList = append(outPutList, "hostid")
	outPutList = append(outPutList, "host")
	interfaceList = append(interfaceList, "interfaceid")
	interfaceList = append(interfaceList, "ip")

	result := ZabbixLogin("Admin", "zabbix@ngcc", "http://192.168.128.181/zabbix/api_jsonrpc.php")
	fmt.Println(result)
	var getHostStruct = response.GetHostModel{
		Jsonrpc: "2.0",
		Method:  "host.get",
		Params: struct {
			Output           []string `json:"output"`
			SelectInterfaces []string `json:"selectInterfaces"`
		}(struct {
			Output           []string
			SelectInterfaces []string
		}{Output: outPutList, SelectInterfaces: interfaceList}),
		Auth: result,
	}
	fmt.Println(getHostStruct)
	r1 = PostProcessingJSON(getHostStruct, "http://192.168.128.181/zabbix/api_jsonrpc.php")
	return r1
}

func main() {
	//ZabbixLogin("Admin", "zabbix@ngcc", "http://192.168.128.181/zabbix/api_jsonrpc.php")
	r1 := GetHostList()
	fmt.Println(r1)
	10099
}
