package routers

import (
	"KnowledgeManage/conf"
	"KnowledgeManage/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session(conf.LoginSessionName).(models.Member)

		if !ok {
			if ctx.Input.IsAjax() {
				jsonData := make(map[string]interface{}, 2)

				jsonData["errcode"] = 403
				jsonData["msg"] = "请登录后再操作"
				returnJSON, _ := json.Marshal(jsonData)

				ctx.ResponseWriter.Write(returnJSON)
			} else {
				ctx.Redirect(302, beego.URLFor("AccountController.Get", "url", ctx.Request.URL.String()))
			}
		}
	}
	beego.InsertFilter("/manage", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/manage/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/list", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/search", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/EditKnowledge/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/CreateKnowledge", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/CreateKnowledge/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/ClassifyManage", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/ClassifyManage/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/BackupData", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/BackupData/*", beego.BeforeRouter, FilterUser)

	var FinishRouter = func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Add("Dreamer-Version", conf.VERSION)
		ctx.ResponseWriter.Header().Add("Dreamer-Site", "http://www.dreamer.cn")
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FinishRouter, false)
}
