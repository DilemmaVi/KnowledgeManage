package controllers

import (
	"KnowledgeManage/models"
	"strings"
	"time"
)

type ClassifyController struct {
	BaseController
}

type ClassifyJsonData struct {
	Msg   string                 `json:"msg"`
	Res   []*models.Classifydata `json:"res"`
	Total int64                  `json:"total"`
}

//Prepare 管理后台准备工作
func (c *ClassifyController) Prepare() {
	c.BaseController.Prepare()

	if !c.Member.IsAdministrator() {
		c.Abort("403")
	}
}

//Get 管理后台页面
func (c *ClassifyController) Get() {
	c.Prepare()
	c.Data["title"] = "分类管理"
	c.Layout = "index.tpl"
	c.TplName = "ClassifyManage.html"
}

//PageLoadJson 获取分页用户信息
func (c *ClassifyController) PageLoadJson() {

	jsonresult := ClassifyJsonData{}
	classifys := &models.Classifydata{}
	classifysmap := make(map[string]string, 3)
	pageSize, _ := c.GetInt("pageSize", 10)
	pageIndex, _ := c.GetInt("pageIndex", 1)
	classifys.Yjfl = c.GetString("yjfl")
	classifys.Ejfl = c.GetString("ejfl")
	classifys.Sjfl = c.GetString("sjfl")
	if classifys.Yjfl != "" {
		classifysmap["Yjfl"] = classifys.Yjfl
	}
	if classifys.Ejfl != "" {
		classifysmap["Ejfl"] = classifys.Ejfl
	}
	if classifys.Sjfl != "" {
		classifysmap["Sjfl"] = classifys.Sjfl
	}
	if len(classifysmap) > 0 {
		classifyresult, totalCount, err := classifys.FindByConditions(pageIndex, pageSize, classifysmap)
		if err == nil {
			jsonresult.Msg = "ok"
			jsonresult.Res = classifyresult
			jsonresult.Total = totalCount
		} else {
			jsonresult.Msg = err.Error()
			jsonresult.Res = classifyresult
			jsonresult.Total = totalCount
		}
	} else {
		classifyresult, totalCount, err := classifys.FindToPager(pageIndex, pageSize)
		if err == nil {
			jsonresult.Msg = "ok"
			jsonresult.Res = classifyresult
			jsonresult.Total = totalCount
		} else {
			jsonresult.Msg = err.Error()
			jsonresult.Res = classifyresult
			jsonresult.Total = totalCount
		}

	}

	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//FindClassify 根据id查找分类信息
func (c *ClassifyController) FindClassify() {
	jsonresult := make(map[string]interface{}, 2)
	ClassifyID, _ := c.GetInt("id", 0)

	if ClassifyID <= 0 {
		jsonresult["msg"] = "参数错误"
		jsonresult["res"] = ""
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	Classifys, err := models.NewClassify().Find(ClassifyID)
	if err == nil {
		jsonresult["msg"] = "ok"
		jsonresult["res"] = Classifys
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	jsonresult["msg"] = err.Error()
	jsonresult["res"] = ""
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return
}

//CreateClassify 添加分类.
func (c *ClassifyController) CreateClassify() {
	jsonresult := ClassifyJsonData{}

	yjfl := strings.TrimSpace(c.GetString("yjfl"))
	ejfl := strings.TrimSpace(c.GetString("ejfl"))
	sjfl := strings.TrimSpace(c.GetString("sjfl"))
	status, _ := c.GetInt("status", 0)

	if status != 0 && status != 1 {
		status = 0
	}

	classify := models.NewClassify()

	classify.Yjfl = yjfl
	classify.Ejfl = ejfl
	classify.Sjfl = sjfl
	classify.Creator = c.Member.Name
	classify.Status = status
	classify.Createtime = time.Now().Format("2006-01-02 15:04:05")

	if err := classify.Add(); err != nil {
		jsonresult.Msg = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	jsonresult.Msg = "ok"
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return
}

//EditClassify 编辑分类信息.
func (c *ClassifyController) EditClassify() {
	jsonresult := make(map[string]interface{}, 1)

	yjfl := strings.TrimSpace(c.GetString("change_yjfl"))
	ejfl := strings.TrimSpace(c.GetString("change_ejfl"))
	sjfl := strings.TrimSpace(c.GetString("change_sjfl"))
	id, _ := c.GetInt("change_id")
	status, _ := c.GetInt("change_status", 0)

	if status != 0 && status != 1 {
		status = 0
	}

	Classify := models.NewClassify()
	if _, err := Classify.Find(id); err != nil {
		jsonresult["msg"] = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	Classify.Yjfl = yjfl
	Classify.Ejfl = ejfl
	Classify.Sjfl = sjfl
	Classify.Status = status
	Classify.Modifytime = time.Now().Format("2006-01-02 15:04:05")
	Classify.Reviser = c.Member.Name

	if err := Classify.Update("yjfl", "ejfl", "sjfl", "status", "Modifytime", "Reviser"); err != nil {
		jsonresult["msg"] = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	jsonresult["msg"] = "ok"
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//DeleteClassify 删除分类
func (c *ClassifyController) DeleteClassify() {
	jsonresult := ClassifyJsonData{}
	ClassifyID, _ := c.GetInt("id", 0)

	if ClassifyID <= 0 {
		jsonresult.Msg = "参数错误"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	_, err := models.NewClassify().Find(ClassifyID)

	if err != nil {
		jsonresult.Msg = "分类不存在"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	err = models.NewClassify().Delete(ClassifyID)

	if err != nil {

		jsonresult.Msg = "删除失败,失败原因:" + err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	jsonresult.Msg = "ok"
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return
}
