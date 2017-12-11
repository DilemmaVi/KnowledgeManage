package controllers

import (
	"KnowledgeManage/commands"
	"KnowledgeManage/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type CreateController struct {
	BaseController
}

type KnowledgesJsonData struct {
	Msg   string                  `json:"msg"`
	Res   []*models.Knowledgedata `json:"res"`
	Total int64                   `json:"total"`
}

//Prepare 管理后台准备工作
func (c *CreateController) Prepare() {
	c.BaseController.Prepare()

	if !c.Member.IsAdministrator() {
		c.Abort("403")
	}
}

//Get 知识创建页面
func (c *CreateController) Get() {
	c.Prepare()

	yjfl, _, err := models.NewClassify().FindByClassify("", -1)
	if err == nil {
		c.Data["yjfl"] = yjfl
	} else {
		c.Data["yjfl"] = ""
	}
	c.Data["title"] = "知识创建"
	c.Data["Model"] = ""
	c.Data["Edit"] = "0"

	c.Layout = "index.tpl"
	c.TplName = "CreateKnowledge.html"
}

//Edit 知识修改页面
func (c *CreateController) Edit() {
	c.Prepare()
	id, _ := c.GetInt(":id")

	yjfl, _, err := models.NewClassify().FindByClassify("", -1)
	if err == nil {
		c.Data["yjfl"] = yjfl
	} else {
		c.Data["yjfl"] = ""
	}
	knowledgedata, err := models.NewKnowledge().Find(id)
	if err != nil {
		beego.Error(err.Error())
		c.Abort("500")
	}

	c.Data["Model"] = knowledgedata
	if id == 0 {
		c.Data["title"] = "知识创建"
		c.Data["Edit"] = "0"
	} else {
		c.Data["Edit"] = "1"
		c.Data["title"] = "知识修改"

	}
	c.Layout = "index.tpl"
	c.TplName = "CreateKnowledge.html"
}

//FindClassify 通过上级分类查询下级分类
func (c *CreateController) FindClassify() {
	jsonresult := make(map[string]interface{}, 2)
	classify := strings.TrimSpace(c.GetString("classify"))
	rank, _ := c.GetInt("rank", 0)

	result, _, err := models.NewClassify().FindByClassify(classify, rank)
	if err == nil {
		jsonresult["msg"] = "ok"
		jsonresult["data"] = result
	} else {
		jsonresult["msg"] = err.Error()
		jsonresult["data"] = ""
	}
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//PageLoadJson 获取分页知识
func (c *CreateController) PageLoadJson() {

	jsonresult := KnowledgesJsonData{}
	Knowledgedatas := &models.Knowledgedata{}
	Knowledgedatasmap := make(map[string]string, 3)
	pageSize, _ := c.GetInt("pageSize", 10)
	pageIndex, _ := c.GetInt("pageIndex", 1)
	Knowledgedatas.Yjfl = c.GetString("yjfl")
	Knowledgedatas.Ejfl = c.GetString("ejfl")
	Knowledgedatas.Sjfl = c.GetString("sjfl")
	if Knowledgedatas.Yjfl != "" {
		Knowledgedatasmap["Yjfl"] = Knowledgedatas.Yjfl
	}
	if Knowledgedatas.Ejfl != "" {
		Knowledgedatasmap["Ejfl"] = Knowledgedatas.Ejfl
	}
	if Knowledgedatas.Sjfl != "" {
		Knowledgedatasmap["Sjfl"] = Knowledgedatas.Sjfl
	}
	if len(Knowledgedatasmap) > 0 {
		Knowledgeresult, totalCount, err := Knowledgedatas.FindByConditions(pageIndex, pageSize, Knowledgedatasmap)
		if err == nil {
			jsonresult.Msg = "ok"
			jsonresult.Res = Knowledgeresult
			jsonresult.Total = totalCount
		} else {
			jsonresult.Msg = err.Error()
			jsonresult.Res = Knowledgeresult
			jsonresult.Total = totalCount
		}
	} else {
		Knowledgeresult, totalCount, err := Knowledgedatas.FindToPager(pageIndex, pageSize)
		if err == nil {
			jsonresult.Msg = "ok"
			jsonresult.Res = Knowledgeresult
			jsonresult.Total = totalCount
		} else {
			jsonresult.Msg = err.Error()
			jsonresult.Res = Knowledgeresult
			jsonresult.Total = totalCount
		}

	}

	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//CreateKnowledge 添加知识.
func (c *CreateController) CreateKnowledge() {
	jsonresult := KnowledgesJsonData{}

	Knowledgedatas := models.NewKnowledge()
	Knowledgedatas.Yjfl = strings.TrimSpace(c.GetString("yjfl"))
	Knowledgedatas.Ejfl = strings.TrimSpace(c.GetString("ejfl"))
	Knowledgedatas.Sjfl = strings.TrimSpace(c.GetString("sjfl"))
	Knowledgedatas.Title = strings.TrimSpace(c.GetString("title"))
	Knowledgedatas.Keyword = strings.TrimSpace(c.GetString("keyword"))
	Knowledgedatas.Content = c.GetString("content")
	Knowledgedatas.Contenthtml = c.GetString("contenthtml")

	Knowledgedatas.Creator = c.Member.Name
	Knowledgedatas.Createtime = time.Now().Format("2006-01-02 15:04:05")

	if err := Knowledgedatas.Add(); err != nil {
		jsonresult.Msg = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	jsonresult.Msg = "ok"
	go commands.IndexKnowledage()
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return
}

//EditKnowledge 编辑知识.
func (c *CreateController) EditKnowledge() {
	jsonresult := make(map[string]interface{}, 1)

	Knowledgedatas := models.NewKnowledge()
	id, _ := c.GetInt("change_id")
	Knowledgedatas.Id = id
	if _, err := Knowledgedatas.Find(id); err != nil {
		jsonresult["msg"] = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	Knowledgedatas.Yjfl = strings.TrimSpace(c.GetString("yjfl"))
	Knowledgedatas.Ejfl = strings.TrimSpace(c.GetString("ejfl"))
	Knowledgedatas.Sjfl = strings.TrimSpace(c.GetString("sjfl"))
	Knowledgedatas.Title = strings.TrimSpace(c.GetString("title"))
	Knowledgedatas.Keyword = strings.TrimSpace(c.GetString("keyword"))
	Knowledgedatas.Content = c.GetString("content")
	Knowledgedatas.Contenthtml = c.GetString("contenthtml")
	Knowledgedatas.Modifytime = time.Now().Format("2006-01-02 15:04:05")
	Knowledgedatas.Reviser = c.Member.Name

	if err := Knowledgedatas.Update("yjfl", "ejfl", "sjfl", "Title", "Keyword", "Content", "Contenthtml", "Modifytime", "Reviser"); err != nil {
		jsonresult["msg"] = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	jsonresult["msg"] = "ok"
	go commands.IndexKnowledage()
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//DeleteKnowledge 删除知识
func (c *CreateController) DeleteKnowledge() {
	jsonresult := KnowledgesJsonData{}
	KnowledgeID, _ := c.GetInt("id", 0)

	if KnowledgeID <= 0 {
		jsonresult.Msg = "参数错误"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	_, err := models.NewKnowledge().Find(KnowledgeID)

	if err != nil {
		jsonresult.Msg = "分类不存在"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	err = models.NewKnowledge().Delete(KnowledgeID)

	if err != nil {

		jsonresult.Msg = "删除失败,失败原因:" + err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	jsonresult.Msg = "ok"
	go commands.IndexKnowledage()
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return
}
