package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/api"
	"github.com/gin-gonic/gin"
)

type RegisterRouter struct {
}

func (s *RegisterRouter) InitRegisterRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.RegisterApi
	{
		plugRouter.POST("", plugApi.ApiName)
	}
}
