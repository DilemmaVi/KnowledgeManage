package models

import (
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
	Reviser    string `json:"revisor"`
}

// TableName 获取对应数据库表名.
func (data *Classifydata) TableName() string {
	return "classifydata"
}

func (data *Classifydata) InsertOrUpdate() error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	if data.Id > 0 {
		_, err := o.Update(data)
		return err
	} else {
		_, err := o.Insert(data)
		return err
	}

}

//分页查找分类.
func (data *Classifydata) FindToPager(pageIndex, pageSize int) ([]*Classifydata, int64, error) {
	o := orm.NewOrm()

	var classifydatas []*Classifydata

	offset := (pageIndex - 1) * pageSize

	totalCount, err := o.QueryTable(data.TableName()).Count()

	if err != nil {
		return classifydatas, 0, err
	}

	_, err = o.QueryTable(data.TableName()).OrderBy("-id").Offset(offset).Limit(pageSize).All(&classifydatas)

	if err != nil {
		return classifydatas, 0, err
	}

	return classifydatas, totalCount, nil
}

//删除一条分类.
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
