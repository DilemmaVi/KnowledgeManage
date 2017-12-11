package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type Knowledgedata struct {
	Id          int    `json:"id"`
	Yjfl        string `json:"yjfl"`
	Ejfl        string `json:"ejfl"`
	Sjfl        string `json:"sjfl"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Contenthtml string `json:"contentHTML"`
	Keyword     string `json:"keyword"`
	Createtime  string `json:"createtime"`
	Modifytime  string `json:"modifytime"`
	Creator     string `json:"creator"`
	Reviser     string `json:"revisor"`
}

// TableName 获取对应数据库表名.
func (data *Knowledgedata) TableName() string {
	return "knowledgedata"
}

func NewKnowledge() *Knowledgedata {
	return &Knowledgedata{}
}

// Add 添加一个知识.
func (data *Knowledgedata) Add() error {
	o := orm.NewOrm()

	if c, err := o.QueryTable(data.TableName()).Filter("Title", data.Title).Count(); err == nil && c > 0 {
		return errors.New("已存在此知识")
	}

	_, err := o.Insert(data)

	if err != nil {
		return err
	}

	return nil
}

// Update 更新知识信息.
func (data *Knowledgedata) Update(cols ...string) error {
	o := orm.NewOrm()

	if _, err := o.Update(data, cols...); err != nil {
		return err
	}

	return nil
}

//Find 通过id查找知识
func (data *Knowledgedata) Find(id int) (*Knowledgedata, error) {
	o := orm.NewOrm()

	data.Id = id
	if err := o.Read(data); err != nil {
		return data, err
	}
	return data, nil
}

//GetAllKnowledgeData 获取所有知识
func (data Knowledgedata) GetAllKnowledgeData() ([]Knowledgedata, error) {
	o := orm.NewOrm()
	var knowledgedatas []Knowledgedata
	if _, err := o.QueryTable(data.TableName()).All(&knowledgedatas); err != nil {
		return knowledgedatas, err
	}
	return knowledgedatas, nil
}

//FindByConditions 根据一级分类、二级分类、三级分类查找知识.
func (data *Knowledgedata) FindByConditions(pageIndex int, pageSize int, conditions map[string]string) ([]*Knowledgedata, int64, error) {
	o := orm.NewOrm()
	QueryResult := o.QueryTable(data.TableName())
	var knowledgedatas []*Knowledgedata
	offset := (pageIndex - 1) * pageSize
	for condition := range conditions {
		QueryResult = QueryResult.Filter(condition, conditions[condition])
	}
	totalCount, err := QueryResult.Count()
	if err != nil {
		return knowledgedatas, 0, err
	}
	_, err = QueryResult.OrderBy("-id").Offset(offset).Limit(pageSize).All(&knowledgedatas)
	if err != nil {
		return knowledgedatas, 0, err
	}

	return knowledgedatas, totalCount, err
}

//FindToPager 分页查找知识.
func (data *Knowledgedata) FindToPager(pageIndex, pageSize int) ([]*Knowledgedata, int64, error) {
	o := orm.NewOrm()

	var knowledgedatas []*Knowledgedata

	offset := (pageIndex - 1) * pageSize

	totalCount, err := o.QueryTable(data.TableName()).Count()

	if err != nil {
		return knowledgedatas, 0, err
	}

	_, err = o.QueryTable(data.TableName()).OrderBy("-id").Offset(offset).Limit(pageSize).All(&knowledgedatas)

	if err != nil {
		return knowledgedatas, 0, err
	}

	return knowledgedatas, totalCount, nil
}

//Delete 删除一条知识.
func (data *Knowledgedata) Delete(Id int) error {
	o := orm.NewOrm()

	err := o.Begin()

	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM knowledgedata WHERE id = ?", Id).Exec()
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
