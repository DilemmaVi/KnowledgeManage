// package conf 为配置相关.
package conf

import "github.com/astaxie/beego"

// 登录用户的Session名
const LoginSessionName = "LoginSessionName"

const VERSION = "1.0"

const CaptchaSessionName = "__captcha__"

const RegexpPhone = `^1[0-9]{10}$`

const RegexpEmail = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`

//允许用户名中出现点号
const RegexpAccount = `^[a-zA-Z][a-zA-z0-9\.]{2,50}$`

// PageSize 默认分页条数.
const PageSize = 15

// 用户权限
const (
	// 超级管理员.
	MemberSuperRole = 0
	//普通管理员.
	MemberAdminRole = 1
	//普通用户.
	MemberGeneralRole = 2
)

//GetAppKey app_key
func GetAppKey() string {
	return beego.AppConfig.DefaultString("app_key", "godoc")
}
