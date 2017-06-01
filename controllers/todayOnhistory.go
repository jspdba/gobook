package controllers

import (
	"github.com/astaxie/beego"
	"gobook/service"
	"net/url"
	"encoding/json"
	"gobook/models"
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

	//先查询数据库，库里有？返回:入库+返回

	today:=new(models.TodayOnHistory)
	today.Key=date
	today.Content = models.FindByKey(date)


	data:=new(map[string]interface{})

	var err error
	if(today.Content==""){
		u, _ := url.Parse(queryEventUrl)

		q := u.Query()
		q.Set("key", key)
		q.Set("date", date)
		b,_:= service.Get(queryEventUrl,&q)

		today.Content = string(b)
		models.ReadOrCreate(today)
	}

	err = json.Unmarshal([]byte(today.Content),data)
	if err!=nil{
		beego.Error(err.Error())
	}

	this.Data["json"] = data
	this.ServeJSON()
}
// @router /todayOnhistory/queryDetail [get]
func(this *TodayOnhistory) QueryDetail() {
	key:=beego.AppConfig.String("API_KEY")
	queryDetailUrl:=beego.AppConfig.String("url.queryDetail")
	e_id:=this.GetString("e_id")


	today:=new(models.TodayOnHistory)
	today.Key=e_id
	today.Content = models.FindByKey(today.Key)


	data:=new(map[string]interface{})

	var err error
	if(today.Content==""){
		u, _ := url.Parse(queryDetailUrl)
		q := u.Query()
		q.Set("key", key)
		q.Set("e_id", e_id)
		b,_ := service.Get(queryDetailUrl,&q)

		today.Content = string(b)
		models.ReadOrCreate(today)
	}

	err = json.Unmarshal([]byte(today.Content),data)
	if err!=nil{
		beego.Error(err.Error())
	}

	this.Data["json"] = data
	this.ServeJSON()
}