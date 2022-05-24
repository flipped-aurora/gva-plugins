package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/tencent_sms/api"
	"github.com/gin-gonic/gin"
)

type TencentSmsRouter struct {
}

func (s *TencentSmsRouter) InitTencentSmsRouter(Router *gin.RouterGroup) {
	smsRouter := Router
	tencentSmsApi := api.ApiGroupApp.TencentSmsApi
	{
		smsRouter.POST("sendTencentSms", tencentSmsApi.SendTencentSms) // 发送测试邮件
	}
}
