package controllers

import (
	"KnowledgeManage/models"

	"github.com/astaxie/beego"
)

type SearchController struct {
	BaseController
}

// Get 搜索首页
func (c *SearchController) Get() {
	c.Prepare()
	var result []map[string]interface{}
	var ejfl interface{}
	yjflList, _, err := models.NewClassify().FindByClassify("", -1)
	if err == nil {
		for _, v := range yjflList {
			first := make(map[string]interface{})
			second := make(map[string]interface{})
			ejflList, _, err := models.NewClassify().FindByClassify(v.(string), 0)
			if err == nil {
				for _, y := range ejflList {
					sjflList, _, err := models.NewClassify().FindByClassify(y.(string), 1)
					if err == nil {
						second[y.(string)] = sjflList
					} else {
						beego.Error(err.Error())
						c.Abort("500")
					}
					ejfl = second
				}
				first[v.(string)] = ejfl
				result = append(result, first)
			} else {
				beego.Error(err.Error())
				c.Abort("500")
			}

		}
	} else {
		beego.Error(err.Error())
		c.Abort("500")
	}

	c.Data["classify"] = result
	c.Data["title"] = "检索"
	c.Layout = "index.tpl"
	c.TplName = "search.html"
}

// Post 搜索首页
func (c *SearchController) Post() {
	c.Prepare()

	var result []map[string]interface{}
	var ejfl interface{}
	yjflList, _, err := models.NewClassify().FindByClassify("", -1)
	if err == nil {
		for _, v := range yjflList {
			first := make(map[string]interface{})
			second := make(map[string]interface{})
			ejflList, _, err := models.NewClassify().FindByClassify(v.(string), 0)
			if err == nil {
				for _, y := range ejflList {
					sjflList, _, err := models.NewClassify().FindByClassify(y.(string), 1)
					if err == nil {
						second[y.(string)] = sjflList
					} else {
						beego.Error(err.Error())
						c.Abort("500")
					}
					ejfl = second
				}
				first[v.(string)] = ejfl
				result = append(result, first)
			} else {
				beego.Error(err.Error())
				c.Abort("500")
			}

		}
	} else {
		beego.Error(err.Error())
		c.Abort("500")
	}

	c.Data["classify"] = result
	c.Data["title"] = "检索"
	c.Data["title"] = "检索结果"
	c.Layout = "index.tpl"
	c.TplName = "result.html"
}
