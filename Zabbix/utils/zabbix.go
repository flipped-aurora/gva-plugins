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

//请求操作+byte转map
func PostProcessingJSON(ProcessStruct interface{}, url string) (ResultMap interface{}, err error) {
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
	result := make(map[string]interface{})
	err = json.Unmarshal(content, &result)
	if err != nil {
		fmt.Println(err)
	}
	//JsonStr := string(content)
	return result["result"], err
}

//Zabbix API登录
func ZabbixLogin() (resultToken interface{}, resulturl string) {
	//获取本地Config.yaml配置的Zabbix参数
	username := global.GlobalConfig.Username
	password := global.GlobalConfig.Password
	url := global.GlobalConfig.Url
	//var resultModel response.ResultModel
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
	resultToken, err := PostProcessingJSON(loginstrcut, url)
	if err != nil {
		return err, resulturl
	}
	resulturl = url
	return resultToken, resulturl
}

//获取所有主机列表
func GetHostList(resultToken interface{}, zabbixurl string) (r1 interface{}) {
	//指定Zabbix返回值参数
	var outPutList []string
	var interfaceList []string
	var groupsList []string
	var templatsList []string
	outPutList = append(outPutList, "name")
	outPutList = append(outPutList, "host")
	outPutList = append(outPutList, "proxy_hostid")
	outPutList = append(outPutList, "status")
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

	r1, err := PostProcessingJSON(getHostStruct, zabbixurl)
	if err != nil {
		fmt.Println(err)
	}
	return r1
}

//根据web状态码进行监控网页状态
func WebMonitor(urls []string) (webCodeStatusList []response.GetWebCodeStatusModel, err error) {
	var getWebCodeStatus response.GetWebCodeStatusModel
	getWebCodeStatusList := make([]response.GetWebCodeStatusModel, 0)
	for _, v := range urls {
		responseGet, _ := http.Get(v)
		getWebCodeStatus.Url = v
		if responseGet.StatusCode != 200 {
			getWebCodeStatus.Status = "该web网址有异常"
			getWebCodeStatusList = append(getWebCodeStatusList, getWebCodeStatus)
		} else {
			getWebCodeStatus.Status = "该web网址运行正常"
			getWebCodeStatusList = append(getWebCodeStatusList, getWebCodeStatus)
		}
	}

	return getWebCodeStatusList, nil
}
