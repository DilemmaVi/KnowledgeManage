package controllers

type ContentController struct {
	BaseController
}

func (c *ContentController) Get() {
	c.Prepare()
	c.Data["title"] = "知识内容"
	c.Layout = "index.tpl"
	c.TplName = "content.html"
}
