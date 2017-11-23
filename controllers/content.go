package controllers

import (
	"KnowledgeManage/models"
	"strings"

	"github.com/astaxie/beego"
)

type ContentController struct {
	BaseController
}

//Index 知识内容页.
func (c *ContentController) Index() {
	id, _ := c.GetInt(":id")

	c.Prepare()
	knowledgedata, err := models.NewKnowledge().Find(id)
	if err == nil {
		c.Data["title"] = knowledgedata.Title
		c.Data["Model"] = knowledgedata
		c.Data["keyword"] = strings.Split(knowledgedata.Keyword, ",")
	} else {
		beego.Error(err)
		c.Abort("500")
	}

	c.Layout = "index.tpl"
	c.TplName = "content.html"
}
