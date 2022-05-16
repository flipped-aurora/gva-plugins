package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/model/response"
)

func GetAlertTiggerList(resultToken interface{}, zabbixurl string) (r1 interface{}, err error) {
	//指定Zabbix返回值参数
	paramsOutput := make([]string, 0)
	paramsOutput = append(paramsOutput, "triggerid")
	paramsOutput = append(paramsOutput, "description")
	paramsOutput = append(paramsOutput, "priority")
	paramsOutput = append(paramsOutput, "status")
	paramsOutput = append(paramsOutput, "selectHosts")
	var AlertParams = response.Params{
		Output:      paramsOutput,
		Filter:      response.Filter(struct{ Value int }{Value: 1}),
		Sortfield:   "priority",
		Sortorder:   "DESC",
		SelectHosts: "extend",
	}

	var getTiggerStruct = response.AlertTiggerModel{
		Jsonrpc: "2.0",
		Method:  "trigger.get",
		Params:  AlertParams,
		Auth:    resultToken,
	}
	r1, err = PostProcessingJSON(getTiggerStruct, zabbixurl)
	if err != nil {
		fmt.Println(err)
	}
	return r1, err
}
