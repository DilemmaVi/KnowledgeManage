// Package models .
package models

import (
	"errors"
	"regexp"
	"strings"

	"KnowledgeManage/conf"
	"KnowledgeManage/utils"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Member struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`

	Role          int    `json:"role"` //用户角色：0 超级管理员 /1 管理员/ 2 普通用户
	RoleName      string `json:"role_name"`
	Status        int    `json:"status"` //用户状态：0 正常/1 禁用
	CreateTime    string `json:"create_time"`
	LastLoginTime string `json:"last_login_time"`
}

// TableName 获取对应数据库表名.
func (m *Member) TableName() string {
	return "members"
}

func NewMember() *Member {
	return &Member{}
}

// Login 用户登录.
func (m *Member) Login(account string, password string) (*Member, error) {
	o := orm.NewOrm()

	member := &Member{}

	err := o.QueryTable(m.TableName()).Filter("account", account).Filter("status", 0).One(member)

	if err != nil {

		logs.Error("用户登录 => ", err)
		return member, ErrMemberNoExist
	}

	ok, err := utils.PasswordVerify(member.Password, password)
	if ok && err == nil {
		m.ResolveRoleName()
		return member, nil

	}
	return member, ErrorMemberPasswordError
}

// Add 添加一个用户.
func (m *Member) Add() error {
	o := orm.NewOrm()

	if c, err := o.QueryTable(m.TableName()).Filter("email", m.Email).Count(); err == nil && c > 0 {
		return errors.New("邮箱已被使用")
	}

	hash, err := utils.PasswordHash(m.Password)

	if err != nil {
		return err
	}

	m.Password = hash

	_, err = o.Insert(m)

	if err != nil {
		return err
	}
	m.ResolveRoleName()
	return nil
}

// Update 更新用户信息.
func (m *Member) Update(cols ...string) error {
	o := orm.NewOrm()

	if m.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if m.Password != "" {
		hash, err := utils.PasswordHash(m.Password)

		if err != nil {
			return err
		}

		m.Password = hash
	}
	if _, err := o.Update(m, cols...); err != nil {
		return err
	}
	return nil
}

func (m *Member) Find(id int) (*Member, error) {
	o := orm.NewOrm()

	m.Id = id
	if err := o.Read(m); err != nil {
		return m, err
	}
	m.ResolveRoleName()
	return m, nil
}

func (m *Member) ResolveRoleName() {
	if m.Role == conf.MemberSuperRole {
		m.RoleName = "超级管理员"
	} else if m.Role == conf.MemberAdminRole {
		m.RoleName = "管理员"
	} else if m.Role == conf.MemberGeneralRole {
		m.RoleName = "普通用户"
	}
}

//FindByAccount 根据账号查找用户
func (m *Member) FindByAccount(account string) (*Member, error) {
	o := orm.NewOrm()

	err := o.QueryTable(m.TableName()).Filter("account", account).One(m)

	if err == nil {
		m.ResolveRoleName()
	}
	return m, err
}

//FindByConditions 根据账号、电话、邮箱查找用户.
func (m *Member) FindByConditions(pageIndex int, pageSize int, conditions map[string]string) ([]*Member, int64, error) {
	o := orm.NewOrm()
	data := o.QueryTable(m.TableName())
	var members []*Member

	offset := (pageIndex - 1) * pageSize
	for condition := range conditions {
		data = data.Filter(condition, conditions[condition])
	}
	totalCount, err := data.Count()
	if err != nil {
		return members, 0, err
	}
	_, err = data.OrderBy("-id").Offset(offset).Limit(pageSize).All(&members)
	if err != nil {
		return members, 0, err
	}
	for _, m := range members {
		m.ResolveRoleName()
	}
	return members, totalCount, err
}

//分页查找用户.
func (m *Member) FindToPager(pageIndex, pageSize int) ([]*Member, int64, error) {
	o := orm.NewOrm()

	var members []*Member

	offset := (pageIndex - 1) * pageSize

	totalCount, err := o.QueryTable(m.TableName()).Count()

	if err != nil {
		return members, 0, err
	}

	_, err = o.QueryTable(m.TableName()).OrderBy("-id").Offset(offset).Limit(pageSize).All(&members)

	if err != nil {
		return members, 0, err
	}

	for _, m := range members {
		m.ResolveRoleName()
	}
	return members, totalCount, nil
}

//检测是否为管理员

func (m *Member) IsAdministrator() bool {
	if m == nil || m.Id <= 0 {
		return false
	}
	return m.Role == 0 || m.Role == 1
}

//根据指定字段查找用户.
func (m *Member) FindByFieldFirst(field string, value interface{}) (*Member, error) {
	o := orm.NewOrm()

	err := o.QueryTable(m.TableName()).Filter(field, value).OrderBy("-id").One(m)

	return m, err
}

//校验用户.
func (m *Member) Valid(is_hash_password bool) error {

	//邮箱不能为空
	if m.Email == "" {
		return ErrMemberEmailEmpty
	}
	if m.Role != conf.MemberGeneralRole && m.Role != conf.MemberSuperRole && m.Role != conf.MemberAdminRole {
		return ErrMemberRoleError
	}
	if m.Status != 0 && m.Status != 1 {
		m.Status = 0
	}
	//邮箱格式校验
	if ok, err := regexp.MatchString(conf.RegexpEmail, m.Email); !ok || err != nil || m.Email == "" {
		return ErrMemberEmailFormatError
	}
	//如果是未加密密码，需要校验密码格式
	if !is_hash_password {
		if l := strings.Count(m.Password, ""); m.Password == "" || l > 50 || l < 6 {
			return ErrMemberPasswordFormatError
		}
	}
	//校验邮箱是否被使用
	if member, err := NewMember().FindByFieldFirst("email", m.Account); err == nil && member.Id > 0 {
		if m.Id > 0 && m.Id != member.Id {
			return ErrMemberEmailExist
		}
		if m.Id <= 0 {
			return ErrMemberEmailExist
		}
	}

	if m.Id > 0 {
		//校验用户是否存在
		if _, err := NewMember().Find(m.Id); err != nil {
			return err
		}
	} else {
		//校验账号格式是否正确
		if ok, err := regexp.MatchString(conf.RegexpAccount, m.Account); m.Account == "" || !ok || err != nil {
			return ErrMemberAccountFormatError
		}
		//校验账号是否被使用
		if member, err := NewMember().FindByAccount(m.Account); err == nil && member.Id > 0 {
			return ErrMemberExist
		}
	}

	return nil
}

//删除一个用户.
func (m *Member) Delete(memberID int) error {
	o := orm.NewOrm()

	err := o.Begin()

	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM members WHERE id = ?", memberID).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	if err = o.Commit(); err != nil {
		o.Rollback()
		return err
	}
	return nil
}
