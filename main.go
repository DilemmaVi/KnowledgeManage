package main

import (
	"KnowledgeManage/commands"
	"KnowledgeManage/controllers"
	_ "KnowledgeManage/routers"

	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	commands.Register()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()

}
