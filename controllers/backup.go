package controllers

import (
	"github.com/astaxie/beego"
)

type BackupController struct {
	beego.Controller
}

func (c *BackupController) Get() {
	c.Data["title"] = "备份知识"
	c.Layout = "index.tpl"
	c.TplName = "BackupData.html"
}
