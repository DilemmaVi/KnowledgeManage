package controllers

import (
	"KnowledgeManage/models"
	"github.com/astaxie/beego"
	"time"
)

type CreateController struct {
	beego.Controller
}

type JsonResult struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func (c *CreateController) Get() {
	c.Data["title"] = "知识创建"
	c.Layout = "index.tpl"
	c.TplName = "CreateKnowledge.html"
}

func (c *CreateController) Post() {
	jsonresult := JsonResult{}
	data := &models.Knowledgedata{}
	Id, _ := c.GetInt("Id")
	data.Id = Id
	data.Yjfl = c.GetString("Yjfl")
	data.Ejfl = c.GetString("Ejfl")
	data.Sjfl = c.GetString("Sjfl")
	data.Title = c.GetString("Title")
	data.Content = c.GetString("Content")
	data.Contenthtml = c.GetString("Contenthtml")
	data.Keyword = c.GetString("Keyword")
	data.Createtime = c.GetString("Createtime")
	data.Modifytime = time.Now().Format("2006-01-02 15:04:05")
	data.Creator = c.GetString("Creator")
	data.Reviser = c.GetString("Reviser")
	err := data.InsertOrUpdate()
	if err == nil {

		jsonresult.Code = 200
		jsonresult.Reason = "ok"
		c.Data["json"] = jsonresult
	} else {

		jsonresult.Code = 500
		jsonresult.Reason = err.Error()
		c.Data["json"] = jsonresult
		c.Ctx.Output.SetStatus(500)

	}

	c.ServeJSON()
	return
}
