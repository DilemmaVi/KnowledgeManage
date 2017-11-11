package controllers

import (
	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	c.Data["title"] = "知识列表"
	c.Layout = "index.tpl"
	c.TplName = "list.html"
}
