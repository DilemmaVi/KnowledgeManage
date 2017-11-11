package routers

import (
	"KnowledgeManage/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.SearchController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/ClassifyManage", &controllers.ClassifyController{})
	beego.Router("/CreateKnowledge", &controllers.CreateController{})
	beego.Router("/BackupData", &controllers.BackupController{})
	beego.Router("/content", &controllers.ContentController{})
	beego.Router("/login", &controllers.HomeController{})
}
