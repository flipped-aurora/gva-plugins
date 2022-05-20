package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/tencent_sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/tencent_sms/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TencentSmsApi struct{}

// SendTencentSms
// @Tags System
// @Summary 发送（腾讯）短信
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /tencentSms/sendTencentSms[post]
func (s *TencentSmsApi) SendTencentSms(c *gin.Context) {
	var req model.TencentModel
	_ = c.ShouldBindJSON(&req)
	if err := service.ServiceGroupApp.SendSms(req.TemplateId, req.Phones, req.TemplateParam); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
