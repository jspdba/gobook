package controllers

import (
	"github.com/astaxie/beego"
	"gobook/models"
	"strings"
	"gobook/wxapp"
)

type WxAppController struct {
	beego.Controller
}

func (this *WxAppController) URLMapping() {
	this.Mapping("/wxapp/userInfo", this.UserInfo)
	this.Mapping("/wxapp/login", this.Login)
}

// @router /wxapp/userInfo [get]
func (this *WxAppController) UserInfo() {
	code := this.GetString("code")
	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"error",
	}
	if(code!=""){
		wxSession,err := wxapp.GetWxSessionKey(code)
		if err==nil{
			//只保存在服务端
			wxSession.SessionKey=""
			jsonMap["code"]=0
			jsonMap["msg"]=""
			jsonMap["data"]=wxSession
		}else{
			jsonMap["msg"]="获取微信信息失败"
		}
	}
	this.Data["json"]=jsonMap
	this.ServeJSON()
}

// @router /wxapp/userInfo [get]
func (this *WxAppController) Login() {
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		_,job:=models.FindJobById(id)
		this.Data["entity"] = job
	}

	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"error",
	}
	this.Data["json"]=jsonMap
	this.ServeJSON()
}