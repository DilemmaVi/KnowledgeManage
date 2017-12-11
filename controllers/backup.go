package controllers

import (
	"KnowledgeManage/models"
	"strconv"

	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego"
)

type BackupController struct {
	BaseController
}

//Prepare 管理后台准备工作
func (c *BackupController) Prepare() {
	c.BaseController.Prepare()

	if !c.Member.IsAdministrator() {
		c.Abort("403")
	}
}

// Get 备份知识展示页面
func (c *BackupController) Get() {
	c.Prepare()
	c.Data["title"] = "备份知识"
	c.Layout = "index.tpl"
	c.TplName = "BackupData.html"
}

//Post 备份知识主程序
func (c *BackupController) Post() {
	c.Prepare()
	Knowledgedatas, err := models.NewKnowledge().GetAllKnowledgeData()
	if err == nil {
		Newxlsx := excelize.NewFile()
		// Create a new sheet.
		NewIndex := Newxlsx.NewSheet("Sheet1")
		Newxlsx.SetCellValue("Sheet1", "A1", "一级分类")
		Newxlsx.SetCellValue("Sheet1", "B1", "二级分类")
		Newxlsx.SetCellValue("Sheet1", "C1", "三级分类")
		Newxlsx.SetCellValue("Sheet1", "D1", "标题")
		Newxlsx.SetCellValue("Sheet1", "E1", "正文")
		Newxlsx.SetCellValue("Sheet1", "F1", "正文HTML")
		Newxlsx.SetCellValue("Sheet1", "G1", "关键词")
		Newxlsx.SetCellValue("Sheet1", "H1", "创建时间")
		Newxlsx.SetCellValue("Sheet1", "I1", "修改时间")
		Newxlsx.SetCellValue("Sheet1", "J1", "创建人")
		Newxlsx.SetCellValue("Sheet1", "K1", "修改人")
		for index, knowledge := range Knowledgedatas {
			index = index + 2
			Newxlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(index), knowledge.Yjfl)
			Newxlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(index), knowledge.Ejfl)
			Newxlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(index), knowledge.Sjfl)
			Newxlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(index), knowledge.Title)
			Newxlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(index), knowledge.Content)
			Newxlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(index), knowledge.Contenthtml)
			Newxlsx.SetCellValue("Sheet1", "G"+strconv.Itoa(index), knowledge.Keyword)
			Newxlsx.SetCellValue("Sheet1", "H"+strconv.Itoa(index), knowledge.Createtime)
			Newxlsx.SetCellValue("Sheet1", "I"+strconv.Itoa(index), knowledge.Modifytime)
			Newxlsx.SetCellValue("Sheet1", "J"+strconv.Itoa(index), knowledge.Creator)
			Newxlsx.SetCellValue("Sheet1", "K"+strconv.Itoa(index), knowledge.Reviser)
		}
		Newxlsx.SetActiveSheet(NewIndex)
		err1 := Newxlsx.SaveAs("static/backup/result.xlsx")
		if err1 == nil {
			c.Ctx.Output.Download("static/backup/result.xlsx")
			return
		}
		beego.Error(err1.Error())
		c.Abort("500")

	} else {
		beego.Error(err.Error())
		c.Abort("500")
	}

}
