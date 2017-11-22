package controllers

type ListController struct {
	BaseController
}

func (c *ListController) Get() {
	c.Prepare()
	c.Data["title"] = "知识列表"
	c.Layout = "index.tpl"
	c.TplName = "list.html"
}
