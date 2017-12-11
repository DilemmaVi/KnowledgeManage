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
	beego.Router("/ClassifyManage/classifys", &controllers.ClassifyController{}, "post:PageLoadJson")
	beego.Router("/ClassifyManage/add", &controllers.ClassifyController{}, "post:CreateClassify")
	beego.Router("/ClassifyManage/delete", &controllers.ClassifyController{}, "post:DeleteClassify")
	beego.Router("/ClassifyManage/find", &controllers.ClassifyController{}, "post:FindClassify")
	beego.Router("/ClassifyManage/edit", &controllers.ClassifyController{}, "post:EditClassify")

	beego.Router("/CreateKnowledge", &controllers.CreateController{})
	beego.Router("/EditKnowledge/:id", &controllers.CreateController{}, "*:Edit")
	beego.Router("/CreateKnowledge/knowledges", &controllers.CreateController{}, "post:PageLoadJson")
	beego.Router("/CreateKnowledge/add", &controllers.CreateController{}, "post:CreateKnowledge")
	beego.Router("/CreateKnowledge/delete", &controllers.CreateController{}, "post:DeleteKnowledge")
	beego.Router("/CreateKnowledge/edit", &controllers.CreateController{}, "post:EditKnowledge")
	beego.Router("/CreateKnowledge/find", &controllers.CreateController{}, "post:FindClassify")

	beego.Router("/BackupData", &controllers.BackupController{})
	beego.Router("/content/:id", &controllers.ContentController{}, "*:Index")

	beego.Router("/login", &controllers.AccountController{})
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/checklogin", &controllers.AccountController{}, "post:Login")
	beego.Router("/upload", &controllers.UploadController{}, "post:UploadFile")

	beego.Router("/manage", &controllers.ManageController{})
	beego.Router("/manage/members", &controllers.ManageController{}, "post:PageLoadJson")
	beego.Router("/manage/add", &controllers.ManageController{}, "post:CreateMember")
	beego.Router("/manage/delete", &controllers.ManageController{}, "post:DeleteMember")
	beego.Router("/manage/find", &controllers.ManageController{}, "post:FindMember")
	beego.Router("/manage/edit", &controllers.ManageController{}, "post:EditMember")
}
