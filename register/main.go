package register

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/router"
	"github.com/gin-gonic/gin"
)

type RegisterPlugin struct {
}

func CreateRegisterPlug(AuthorityId string) *RegisterPlugin {
	global.GlobalConfig.AuthorityId = AuthorityId
	return &RegisterPlugin{}
}

func (*RegisterPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRegisterRouter(group)
}

func (*RegisterPlugin) RouterPath() string {
	return "register"
}
