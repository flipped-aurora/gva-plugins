package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/router"
	"github.com/gin-gonic/gin"
)

type esPlugin struct{}

//
// CreateEsPlug
//  @Description: 创建es插件
//  @param Host
//  @param Port
//  @param User
//  @param Pass
//  @return *esPlugin
//  @author fyk
//
func CreateEsPlug(Host, Port, User, Pass string) *esPlugin {
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Port = Port
	global.GlobalConfig.User = User
	global.GlobalConfig.Pass = Pass
	return &esPlugin{}
}

func (*esPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEsRouter(group)
}

func (*esPlugin) RouterPath() string {
	return "es"
}
