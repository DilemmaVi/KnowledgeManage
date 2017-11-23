package controllers

import "time"

type UploadController struct {
	BaseController
}

//upload 知识内容页.
func (c *UploadController) UploadFile() {
	jsonresult := make(map[string]interface{}, 2)
	f, _, err := c.GetFile("imageFile")
	defer f.Close()
	if err != nil {
		jsonresult["errno"] = 1
		jsonresult["data"] = ""
	} else {
		c.SaveToFile("imageFile", "./static/img/"+time.Now().Format("2006_01_02_15_04_05")+".jpg")
		jsonresult["errno"] = 0
		jsonresult["data"] = [1]string{"/static/img/" + time.Now().Format("2006_01_02_15_04_05") + ".jpg"}
	}
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return
}
