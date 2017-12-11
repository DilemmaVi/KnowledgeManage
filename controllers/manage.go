package controllers

import (
	"KnowledgeManage/conf"
	"KnowledgeManage/models"
	"regexp"
	"strings"
	"time"
)

type ManageController struct {
	BaseController
}

type MembersJsonData struct {
	Msg   string           `json:"msg"`
	Res   []*models.Member `json:"res"`
	Total int64            `json:"total"`
}

//Prepare 管理后台准备工作
func (c *ManageController) Prepare() {
	c.BaseController.Prepare()

	if !c.Member.IsAdministrator() {
		c.Abort("403")
	}
}

//Get 管理后台页面
func (c *ManageController) Get() {
	c.Prepare()
	c.Data["title"] = "管理后台"
	c.Layout = "index.tpl"
	c.TplName = "manage.html"
}

//PageLoadJson 获取分页用户信息
func (c *ManageController) PageLoadJson() {

	jsonresult := MembersJsonData{}
	members := &models.Member{}
	membersmap := make(map[string]string, 3)
	pageSize, _ := c.GetInt("pageSize", 10)
	pageIndex, _ := c.GetInt("pageIndex", 1)
	members.Name = c.GetString("Name")
	members.Phone = c.GetString("Tel")
	members.Email = c.GetString("Email")
	if members.Name != "" {
		membersmap["Name"] = members.Name
	}
	if members.Phone != "" {
		membersmap["Phone"] = members.Phone
	}
	if members.Email != "" {
		membersmap["Email"] = members.Email
	}
	if len(membersmap) > 0 {
		memberresult, totalCount, err := members.FindByConditions(pageIndex, pageSize, membersmap)
		if err == nil {
			jsonresult.Msg = "ok"
			jsonresult.Res = memberresult
			jsonresult.Total = totalCount
		} else {
			jsonresult.Msg = err.Error()
			jsonresult.Res = memberresult
			jsonresult.Total = totalCount
		}
	} else {
		memberresult, totalCount, err := members.FindToPager(pageIndex, pageSize)
		if err == nil {
			jsonresult.Msg = "ok"
			jsonresult.Res = memberresult
			jsonresult.Total = totalCount
		} else {
			jsonresult.Msg = err.Error()
			jsonresult.Res = memberresult
			jsonresult.Total = totalCount
		}

	}

	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//FindMember 根据id查找用户信息
func (c *ManageController) FindMember() {
	jsonresult := make(map[string]interface{}, 2)
	memberID, _ := c.GetInt("id", 0)

	if memberID <= 0 {
		jsonresult["msg"] = "参数错误"
		jsonresult["res"] = ""
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	members, err := models.NewMember().Find(memberID)
	if err == nil {
		jsonresult["msg"] = "ok"
		jsonresult["res"] = members
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

//CreateMember 添加用户.
func (c *ManageController) CreateMember() {
	jsonresult := MembersJsonData{}

	account := strings.TrimSpace(c.GetString("account"))
	name := strings.TrimSpace(c.GetString("name"))
	password := strings.TrimSpace(c.GetString("password"))
	email := strings.TrimSpace(c.GetString("email"))
	phone := strings.TrimSpace(c.GetString("phone"))
	role, _ := c.GetInt("role", 1)
	status, _ := c.GetInt("status", 0)

	// if ok, err := regexp.MatchString(conf.RegexpAccount, account); account == "" || !ok || err != nil {
	// 	jsonresult.Msg = "账号只能由英文字母数字组成，且在3-50个字符" + account
	// 	c.Data["json"] = jsonresult
	// 	c.ServeJSON()
	// 	return
	// }

	if ok, err := regexp.MatchString(conf.RegexpPhone, phone); !ok || err != nil || phone == "" {
		jsonresult.Msg = "手机号码格式不正确"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if l := strings.Count(password, ""); password == "" || l > 50 || l < 6 {
		jsonresult.Msg = "密码必须在6-50个字符之间"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if ok, err := regexp.MatchString(conf.RegexpEmail, email); !ok || err != nil || email == "" {
		jsonresult.Msg = "邮箱格式不正确"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if role != 0 && role != 1 && role != 2 {
		role = 1
	}
	if status != 0 && status != 1 {
		status = 0
	}

	member := models.NewMember()

	if _, err := member.FindByAccount(account); err == nil && member.Id > 0 {
		jsonresult.Msg = "账号已存在"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	member.Account = account
	member.Name = name
	member.Password = password
	member.Role = role
	member.Email = email
	member.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	if phone != "" {
		member.Phone = phone
	}

	if c.Member.Role != 0 && member.Role == 0 {
		jsonresult.Msg = "你无权添加超级管理员"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}
	if err := member.Add(); err != nil {
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

//EditMember 编辑用户信息.
func (c *ManageController) EditMember() {
	jsonresult := make(map[string]interface{}, 1)

	name := strings.TrimSpace(c.GetString("change_name"))
	password := strings.TrimSpace(c.GetString("change_password"))
	email := strings.TrimSpace(c.GetString("change_email"))
	phone := strings.TrimSpace(c.GetString("change_phone"))
	role, _ := c.GetInt("change_role", 1)
	id, _ := c.GetInt("change_id")
	status, _ := c.GetInt("change_status", 0)

	if ok, err := regexp.MatchString(conf.RegexpPhone, phone); !ok || err != nil || phone == "" {
		jsonresult["msg"] = "手机号码格式不正确"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if ok, err := regexp.MatchString(conf.RegexpEmail, email); !ok || err != nil || email == "" {
		jsonresult["msg"] = "邮箱格式不正确"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if role != 0 && role != 1 && role != 2 {
		role = 1
	}
	if status != 0 && status != 1 {
		status = 0
	}

	member := models.NewMember()

	m, err := member.Find(id)

	if err != nil {
		jsonresult["msg"] = err.Error()
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if c.Member.Role != 0 && m.Role == 0 {

		jsonresult["msg"] = "你无权修改超级管理员"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	member.Id = id
	member.Name = name
	member.Password = password
	member.Role = role
	member.Email = email
	member.Status = status
	if phone != "" {
		member.Phone = phone
	}

	if member.Password != "" {
		if l := strings.Count(password, ""); password == "" || l > 50 || l < 6 {
			jsonresult["msg"] = "密码必须在6-50个字符之间"
			c.Data["json"] = jsonresult
			c.ServeJSON()
			return
		}
		if err := member.Update("name", "password", "role", "email", "phone", "status"); err != nil {
			jsonresult["msg"] = err.Error()
			c.Data["json"] = jsonresult
			c.ServeJSON()
			return
		}
	} else {
		if err := member.Update("name", "role", "email", "phone", "status"); err != nil {
			jsonresult["msg"] = err.Error()
			c.Data["json"] = jsonresult
			c.ServeJSON()
			return
		}
	}
	jsonresult["msg"] = "ok"
	c.Data["json"] = jsonresult
	c.ServeJSON()
	return

}

//DeleteMember 删除用户
func (c *ManageController) DeleteMember() {
	jsonresult := MembersJsonData{}
	memberID, _ := c.GetInt("id", 0)

	if memberID <= 0 {
		jsonresult.Msg = "参数错误"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	m, err := models.NewMember().Find(memberID)

	if c.Member.Role != 0 && m.Role == 0 {

		jsonresult.Msg = "你无权删除超级管理员"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	if err != nil {
		jsonresult.Msg = "用户不存在"
		c.Data["json"] = jsonresult
		c.ServeJSON()
		return
	}

	err = models.NewMember().Delete(memberID)

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
