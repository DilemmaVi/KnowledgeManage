package controllers

import (
	"github.com/astaxie/beego"
)

type ClassifyController struct {
	beego.Controller
}

func (c *ClassifyController) Get() {
	c.Data["title"] = "分类管理"
	c.Layout = "index.tpl"
	c.TplName = "ClassifyManage.html"
}
