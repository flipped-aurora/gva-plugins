package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/api"
	"github.com/gin-gonic/gin"
)

type ZabbixRouter struct{}

func (s *ZabbixRouter) InitZabbixRouter(Router *gin.RouterGroup) {
	zabbixRouter := Router.Use(middleware.OperationRecord())
	GetHostAll := api.ApiGroupApp.GetHostAll
	GetWebStatusCode := api.ApiGroupApp.GetWebMonitorStatus
	GetTiggerList := api.ApiGroupApp.GetTiggerList
	{
		zabbixRouter.GET("getHostAll", GetHostAll)        // 获取主机列表
		zabbixRouter.GET("getTiggerList", GetTiggerList)  //获取正在告警列表
		zabbixRouter.POST("getwebcode", GetWebStatusCode) // 监控Web网页状态
	}
}
