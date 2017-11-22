package controllers

import (
	"KnowledgeManage/conf"
	"KnowledgeManage/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Member *models.Member
}

// Prepare 预处理.
// 控制器预处理程序
func (c *BaseController) Prepare() {
	c.Data["Member"] = ""

	if member, ok := c.GetSession(conf.LoginSessionName).(models.Member); ok && member.Id > 0 {
		c.Member = &member
		c.Data["Member"] = c.Member.Name
	}
	c.Data["BaseUrl"] = c.Ctx.Input.Scheme() + "://" + c.Ctx.Request.Host
}

// SetMember 获取或设置当前登录用户信息,如果 Id 小于 0 则标识删除 Session
func (c *BaseController) SetMember(member models.Member) {

	if member.Id <= 0 {
		c.DelSession(conf.LoginSessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {
		c.SetSession(conf.LoginSessionName, member)
		c.SetSession("uid", member.Id)
	}
}

//ShowErrorPage 显示错误信息页面.
func (c *BaseController) ShowErrorPage(errCode int, errMsg string) {
	c.TplName = "errors/error.tpl"
	c.Data["ErrorMessage"] = errMsg
	c.Data["ErrorCode"] = errCode
	c.StopRun()
}
