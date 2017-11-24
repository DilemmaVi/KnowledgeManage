package controllers

import (
	"KnowledgeManage/models"
	"KnowledgeManage/utils"
	"math"

	"github.com/astaxie/beego"
)

type ListController struct {
	BaseController
}

func (c *ListController) Get() {

	pageIndex, _ := c.GetInt("page", 1)
	cOne := c.GetString("yjfl")
	cTwo := c.GetString("ejfl")
	cThree := c.GetString("sjfl")
	classifys := make(map[string]string, 3)

	if cOne != "" {
		classifys["yjfl"] = cOne
	}
	if cTwo != "" {
		classifys["ejfl"] = cTwo
	}
	if cThree != "" {
		classifys["sjfl"] = cThree
	}

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

	knowledgedatas, totalCount, err := models.NewKnowledge().FindByConditions(pageIndex, 10, classifys)

	if err != nil {
		beego.Error(err.Error())
		c.Abort("500")
	}

	if totalCount > 0 {
		html := utils.GetPagerHtml(c.Ctx.Request.RequestURI, pageIndex, 10, int(totalCount))

		c.Data["PageHtml"] = html
	} else {
		c.Data["PageHtml"] = ""
	}

	c.Data["TotalPages"] = int(math.Ceil(float64(totalCount) / float64(10)))
	c.Data["knowledgedatas"] = knowledgedatas
	c.Data["classify"] = result
	c.Data["title"] = "知识列表"
	c.Layout = "index.tpl"
	c.TplName = "list.html"
}
