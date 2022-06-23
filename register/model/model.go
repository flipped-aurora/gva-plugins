package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	uuid "github.com/satori/go.uuid"
)

type Request struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type Response struct {
	global.GVA_MODEL
	UUID        uuid.UUID             `json:"uuid" gorm:"comment:用户UUID"`                                                           // 用户UUID
	Username    string                `json:"userName" gorm:"comment:用户登录名"`                                                        // 用户登录名
	Password    string                `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode    string                `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string                `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string                `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId string                `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority   system.SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []system.SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string                `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email       string                `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}
