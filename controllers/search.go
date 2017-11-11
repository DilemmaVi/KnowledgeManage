package controllers

import (
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	c.Data["title"] = "检索"
	c.Layout = "index.tpl"
	c.TplName = "search.html"
}
