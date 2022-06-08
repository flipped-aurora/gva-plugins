package api

import (
	systemApi "github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterApi struct{}

// @Tags Register
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /register/routerName[post]
func (p *RegisterApi) ApiName(c *gin.Context) {
	var plug model.Request
	_ = c.ShouldBindJSON(&plug)
	if res, err := service.ServiceGroupApp.PlugService(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		var baseApi systemApi.BaseApi
		baseApi.TokenNext(c, *res)
	}
}
