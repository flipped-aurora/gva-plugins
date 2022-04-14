package service

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/utils"
	"strings"
)

type ZabbixService struct{}

func GetHost() (err error, HostAll interface{}) {
	r1, r2 := utils.ZabbixLogin()
	HostAll = utils.GetHostList(r1, r2)
	return
}

func GetAlertTiggerList() (TiggersAll interface{}, err error) {
	r1, r2 := utils.ZabbixLogin()
	TiggersAll, err = utils.GetAlertTiggerList(r1, r2)
	return TiggersAll, err
}

func WebMonitor(urls string) (Code interface{}, err error) {
	urlsList := strings.Split(urls, ",")
	Code123, err := utils.WebMonitor(urlsList)
	if err != nil {
		fmt.Println(err)
	}
	return Code123, err
}
