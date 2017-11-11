package main

import (
	"KnowledgeManage/commands"
	_ "KnowledgeManage/routers"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	commands.Register()
	beego.Run()
}
