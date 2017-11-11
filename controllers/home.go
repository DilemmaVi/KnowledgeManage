package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["title"] = "登录"
	c.Layout = "index.tpl"
	c.TplName = "login.html"
}
