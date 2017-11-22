package commands

import (
	"KnowledgeManage/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func Register() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(new(models.Knowledgedata))
	orm.RegisterModel(new(models.Member))
	orm.RegisterModel(new(models.Classifydata))
	createTable()
}

func createTable() {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("CREATE TABLE KnowledgeData (Id integer not null primary key, YJFL text,EJFL text,SJFL text,Title text,Content text,ContentHTML text,KeyWord text,CreateTime text,ModifyTime text,Creator text,Reviser text);").Exec()
	o.Raw("CREATE TABLE members (Id integer not null primary key, account text,password text,name text,email text,phone text,role text,role_name text,status text,create_time text,last_login_time text);").Exec()
	o.Raw("CREATE TABLE Classifydata (Id integer not null primary key, Yjfl text,Ejfl text,Sjfl text,Status,CreateTime text,ModifyTime text,Creator text,Reviser text);").Exec()

}
