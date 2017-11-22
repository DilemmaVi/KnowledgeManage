package controllers

import (
	"time"

	"KnowledgeManage/conf"
	"KnowledgeManage/models"
	"KnowledgeManage/utils"

	"github.com/astaxie/beego"
)

type AccountController struct {
	BaseController
}

func (c *AccountController) Get() {
	c.Prepare()
	var remember struct {
		MemberId int
		Account  string
		Time     time.Time
	}
	//如果Cookie中存在登录信息
	if cookie, ok := c.GetSecureCookie(conf.GetAppKey(), "login"); ok {

		if err := utils.Decode(cookie, &remember); err == nil {
			if member, err := models.NewMember().Find(remember.MemberId); err == nil {
				c.SetMember(*member)
				c.Redirect(beego.URLFor("ListController.Get"), 302)
				c.StopRun()
			}
		}
	}
	c.Data["title"] = "登录"
	c.Layout = "index.tpl"
	c.TplName = "login.html"
}

// Login 用户登录.
func (c *AccountController) Login() {
	c.Prepare()
	var remember struct {
		MemberId int
		Account  string
		Time     time.Time
	}
	var jsonresult = make(map[string]string, 1)

	if c.Ctx.Input.IsPost() {
		account := c.GetString("account")
		password := c.GetString("password")

		member, err := models.NewMember().Login(account, password)

		//如果没有数据
		if err == nil {
			member.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")
			member.Update("LastLoginTime")

			c.SetMember(*member)

			remember.MemberId = member.Id
			remember.Account = member.Account
			remember.Time = time.Now()
			v, err := utils.Encode(remember)
			if err == nil {
				c.SetSecureCookie(conf.GetAppKey(), "login", v)
			}

			jsonresult["msg"] = "ok"

		} else {

			jsonresult["msg"] = "账号或密码错误"

		}

		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
}
