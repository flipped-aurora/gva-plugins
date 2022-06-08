package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	plugGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/register/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/model"
	userService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/mojocn/base64Captcha"
	uuid "github.com/satori/go.uuid"
)

type RegisterService struct{}

func (e *RegisterService) PlugService(req model.Request) (res *system.SysUser, err error) {
	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		// TODO resopne need
		return res, err
	}
	var (
		store = base64Captcha.DefaultMemStore
		user  system.SysUser
		us    *userService.UserService
	)
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		return res, errors.New("验证码错误")
	}
	u := &system.SysUser{Username: req.Username, Password: req.Password}
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		return res, errors.New("用户名已注册")
	}
	if user.Username != "" && user.Password != "" {
		return res, errors.New("用户名已注册")
	}

	var sysAuthority systemReq.Register
	sysAuthority.Username = u.Username
	sysAuthority.NickName = u.NickName
	sysAuthority.Password = u.Password
	// TODO 这个 1 是角色ID
	sysAuthority.AuthorityId = plugGlobal.GlobalConfig.AuthorityId
	sysAuthority.AuthorityIds = append(sysAuthority.AuthorityIds, "1")
	// 因为上面定义过，且得到了数据库默认的值，所以直接使用
	user.Password = u.Password
	user.UUID = uuid.NewV4()
	user.Username = u.Username
	user.NickName = u.Username
	user.AuthorityId = plugGlobal.GlobalConfig.AuthorityId

	for _, v := range sysAuthority.AuthorityIds {
		user.Authorities = append(user.Authorities, system.SysAuthority{
			AuthorityId: v,
			//DefaultRouter: "dashboard", //TODO 疑问，系统注册的时候有这个参数
		})
	}

	if rest, err := us.Register(*u); err != nil {
		return &rest, errors.New("注册失败!")
	}
	if res, err = us.Login(u); err != nil {
		return res, errors.New("登陆失败!")
	}
	return res, nil
	// 前面的代码 拿不到正确的 user，所以需要再次查询一次

}
