package controllers

import "github.com/astaxie/beego"

type WxController struct {
	beego.Controller
}
var (
	AppID = ""
	AppSecret=""
)
func (this *WxController) URLMapping() {
	//this.Mapping("/wx/url/info", this.UrlInfo)
	//this.Mapping("/wx/search", this.Search)
}
