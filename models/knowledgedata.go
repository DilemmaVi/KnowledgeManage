package models

import (
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
	return "Knowledgedata"
}

func (data *Knowledgedata) InsertOrUpdate() error {
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

//分页查找知识.
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

//删除一条知识.
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
