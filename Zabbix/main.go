package Zabbix

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/router"
	"github.com/gin-gonic/gin"
)

type zabbixPlugin struct{}

func CreateZabbixPlugin(url string, username string, password string) *zabbixPlugin {
	global.GlobalConfig.Url = url
	global.GlobalConfig.Username = username
	global.GlobalConfig.Password = password

	return &zabbixPlugin{}
}

func (*zabbixPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitZabbixRouter(group)
}

func (*zabbixPlugin) RouterPath() string {
	return "zabbix"
}
