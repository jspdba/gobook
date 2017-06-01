package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type TodayOnHistory struct {
	Id int			`orm:"auto"`
	Key string		`orm:"unique;"`
	Content string		`orm:"null;type(text)"`
	CreateDate  time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 	`orm:"auto_now;type(datetime)"`
}

//创建并获取内容
func ReadOrCreate(today *TodayOnHistory) string{
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(today, "Key"); err == nil {
		if created {
			return today.Content
		}
	}
	return ""
}

//只取内容
func FindByKey(key string) string{
	o := orm.NewOrm()
	var today TodayOnHistory
	err := o.QueryTable(today).Filter("Key", key).One(&today)
	if err!=nil{
		return ""
	}
	return today.Content
}

func init() {
	orm.RegisterModel(new(TodayOnHistory))
}