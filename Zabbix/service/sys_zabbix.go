package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/utils"
)

type ZabbixService struct{}

func GetHost() (err error, HostAll string) {
	r1, r2 := utils.ZabbixLogin()
	HostAll = utils.GetHostList(r1, r2)
	return
}
