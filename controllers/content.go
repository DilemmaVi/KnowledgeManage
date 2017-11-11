package controllers

import (
	"github.com/astaxie/beego"
)

type ContentController struct {
	beego.Controller
}

func (c *ContentController) Get() {
	c.Data["title"] = "知识内容"
	c.Layout = "index.tpl"
	c.TplName = "content.html"
}
