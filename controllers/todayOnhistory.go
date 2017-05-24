package controllers

import (
	"github.com/astaxie/beego"
	"gobook/service"
	"net/url"
	"encoding/json"
)

type TodayOnhistory struct {
	beego.Controller
}

func(this *TodayOnhistory) URLMapping() {
	this.Mapping("/todayOnhistory/queryEvent", this.QueryEvent)
	this.Mapping("/todayOnhistory/queryDetail", this.QueryDetail)
}

// @router /todayOnhistory/queryEvent [get]
func(this *TodayOnhistory) QueryEvent() {
	key:=beego.AppConfig.String("API_KEY")
	queryEventUrl:=beego.AppConfig.String("url.queryEvent")
	date:=this.GetString("date")

	u, _ := url.Parse(queryEventUrl)

	q := u.Query()
	q.Set("key", key)
	q.Set("date", date)
	b,_:= service.Get(queryEventUrl,&q)

	data:=new(map[string]interface{})
	err:=json.Unmarshal(b,data)
	if err!=nil{
		beego.Error(err)
	}
	this.Data["json"] = data
	this.ServeJSON()
}
// @router /todayOnhistory/QueryDetail [get]
func(this *TodayOnhistory) QueryDetail() {
	key:=beego.AppConfig.String("API_KEY")
	queryDetailUrl:=beego.AppConfig.String("url.queryDetail")
	e_id:=this.GetString("e_id")

	u, _ := url.Parse(queryDetailUrl)
	q := u.Query()
	q.Set("key", key)
	q.Set("e_id", e_id)
	b,_ := service.Get(queryDetailUrl,&q)
	data:=new(map[string]interface{})
	err:=json.Unmarshal(b,data)
	if err!=nil{
		beego.Error(err)
	}
	this.Data["json"] = data
	this.ServeJSON()
}