package controllers

type BackupController struct {
	BaseController
}

func (c *BackupController) Get() {
	c.Prepare()
	c.Data["title"] = "备份知识"
	c.Layout = "index.tpl"
	c.TplName = "BackupData.html"
}
