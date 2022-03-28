package utils

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/model/response"
	"io/ioutil"
	"net/http"
	"strings"
)

//请求操作+byte转Json
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

//Zabbix API登录
func ZabbixLogin() (resultToken string, resulturl string) {
	//获取本地Config.yaml配置的Zabbix参数
	username := global.GlobalConfig.Username
	password := global.GlobalConfig.Password
	url := global.GlobalConfig.Url
	var resultModel response.ResultModel
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
	//进行POST请求及Json转换
	json.Unmarshal([]byte(PostProcessingJSON(loginstrcut, url)), &resultModel)
	resultToken = resultModel.Result
	resulturl = url
	fmt.Println(resulturl, resultToken)
	return resultToken, resulturl
}

//获取所有主机列表
func GetHostList(resultToken string, zabbixurl string) (r1 string) {
	//指定Zabbix返回值参数
	var outPutList []string
	var interfaceList []string
	var groupsList []string
	var templatsList []string
	outPutList = append(outPutList, "name")
	outPutList = append(outPutList, "host")
	outPutList = append(outPutList, "proxy_hostid")
	interfaceList = append(interfaceList, "interfaceid")
	interfaceList = append(interfaceList, "ip")
	groupsList = append(groupsList, "name")
	templatsList = append(templatsList, "name")

	var getHostStruct = response.GetHostModel{
		Jsonrpc: "2.0",
		Method:  "host.get",
		Params: struct {
			Output                []string `json:"output"`
			SelectInterfaces      []string `json:"selectInterfaces"`
			SelectGroups          []string `json:"selectGroups"`
			SelectParentTemplates []string `json:"selectParentTemplates"`
		}(struct {
			Output                []string
			SelectInterfaces      []string
			SelectGroups          []string
			SelectParentTemplates []string
		}{Output: outPutList, SelectInterfaces: interfaceList, SelectGroups: groupsList, SelectParentTemplates: templatsList}),
		Auth: resultToken,
	}

	r1 = PostProcessingJSON(getHostStruct, zabbixurl)
	return r1
}

func GetItemsAll(resultToken string, zabbixurl string) {

}