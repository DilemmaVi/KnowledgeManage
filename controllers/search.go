package controllers

type SearchController struct {
	BaseController
}

func (c *SearchController) Get() {
	c.Prepare()
	c.Data["title"] = "检索"
	c.Layout = "index.tpl"
	c.TplName = "search.html"
}
