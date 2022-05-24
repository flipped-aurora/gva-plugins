package sms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/ali_sms/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/ali_sms/router"
	"github.com/gin-gonic/gin"
)

type aliSmsPlugin struct {
}

func CreateAliSmsPlug(AccessKeyId, AccessSecret, SignName string) *aliSmsPlugin {
	global.GlobalConfig.AccessKeyId = AccessKeyId
	global.GlobalConfig.AccessSecret = AccessSecret
	global.GlobalConfig.SignName = SignName
	return &aliSmsPlugin{}
}

func (*aliSmsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitAliSmsRouter(group)
}

func (*aliSmsPlugin) RouterPath() string {
	return "aliSms"
}
