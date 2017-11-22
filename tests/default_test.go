package test

import (
	"KnowledgeManage/models"
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func Register() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(new(models.Knowledgedata))
	orm.RegisterModel(new(models.Member))
	orm.RegisterModel(new(models.Classifydata))
}

func Test_FindByClassify(t *testing.T) {
	Register()
	classify := models.NewClassify()
	data, _, _ := classify.FindByClassify("测试一级", -1)
	fmt.Println(data)
}
