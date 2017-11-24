package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type Classifydata struct {
	Id         int    `json:"id"`
	Yjfl       string `json:"yjfl"`
	Ejfl       string `json:"ejfl"`
	Sjfl       string `json:"sjfl"`
	Createtime string `json:"createtime"`
	Modifytime string `json:"modifytime"`
	Creator    string `json:"creator"`
	Reviser    string `json:"revisor"` //修改人
	Status     int    `json:"status"`  //状态：0 正常/1 禁用
}

// TableName 获取对应数据库表名.
func (data *Classifydata) TableName() string {
	return "classifydata"
}

func NewClassify() *Classifydata {
	return &Classifydata{}
}

// Add 添加一个分类.
func (data *Classifydata) Add() error {
	o := orm.NewOrm()

	if c, err := o.QueryTable(data.TableName()).Filter("Yjfl", data.Yjfl).Filter("Ejfl", data.Ejfl).Filter("Sjfl", data.Sjfl).Count(); err == nil && c > 0 {
		return errors.New("已存在此分类")
	}

	_, err := o.Insert(data)

	if err != nil {
		return err
	}
	return nil
}

// Update 更新分类信息.
func (data *Classifydata) Update(cols ...string) error {
	o := orm.NewOrm()

	if _, err := o.Update(data, cols...); err != nil {
		return err
	}
	return nil
}

//Find 通过id查找分类
func (data *Classifydata) Find(id int) (*Classifydata, error) {
	o := orm.NewOrm()

	data.Id = id
	if err := o.Read(data); err != nil {
		return data, err
	}
	return data, nil
}

//FindByClassify 根据某一分类查找下一级分类
func (data *Classifydata) FindByClassify(classify string, rank int) (orm.ParamsList, int64, error) {
	o := orm.NewOrm()
	var col string
	var resultcol string
	var classfiys orm.ParamsList
	if rank == -1 {
		col = "yjfl"
		resultcol = "yjfl"
	} else if rank == 0 {
		col = "yjfl"
		resultcol = "ejfl"

	} else {
		col = "ejfl"
		resultcol = "sjfl"
	}
	if rank == -1 {
		num, err := o.QueryTable(data.TableName()).GroupBy(resultcol).ValuesFlat(&classfiys, resultcol)
		return classfiys, num, err
	}
	num, err := o.QueryTable(data.TableName()).Filter(col, classify).GroupBy(resultcol).ValuesFlat(&classfiys, resultcol)
	return classfiys, num, err

}

//FindByConditions 根据一级分类、二级分类、三级分类查找分类信息.
func (data *Classifydata) FindByConditions(pageIndex int, pageSize int, conditions map[string]string) ([]*Classifydata, int64, error) {
	o := orm.NewOrm()
	QueryResult := o.QueryTable(data.TableName())
	var classifys []*Classifydata
	offset := (pageIndex - 1) * pageSize
	for condition := range conditions {
		QueryResult = QueryResult.Filter(condition, conditions[condition])
	}
	totalCount, err := QueryResult.Count()
	if err != nil {
		return classifys, 0, err
	}
	_, err = QueryResult.OrderBy("-id").Offset(offset).Limit(pageSize).All(&classifys)
	if err != nil {
		return classifys, 0, err
	}

	return classifys, totalCount, err
}

//FindToPager 分页查找分类.
func (data *Classifydata) FindToPager(pageIndex, pageSize int) ([]*Classifydata, int64, error) {
	o := orm.NewOrm()

	var classifydatas []*Classifydata

	offset := (pageIndex - 1) * pageSize

	totalCount, err := o.QueryTable(data.TableName()).Count()

	if err != nil {
		return classifydatas, 0, err
	}

	_, err = o.QueryTable(data.TableName()).OrderBy("-Yjfl").Offset(offset).Limit(pageSize).All(&classifydatas)

	if err != nil {
		return classifydatas, 0, err
	}

	return classifydatas, totalCount, nil
}

//Delete 删除一条分类.
func (data *Classifydata) Delete(Id int) error {
	o := orm.NewOrm()

	err := o.Begin()

	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM classifydata WHERE id = ?", Id).Exec()
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
