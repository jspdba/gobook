package controllers

import "github.com/astaxie/beego"

type DecodeController struct {
	beego.Controller
}
func (this *DecodeController) URLMapping() {
	this.Mapping("/str/decode", this.Decode)
}
//@router /str/decode [get]
func (this *DecodeController) Decode() {
	this.TplName="str/decode.tpl"
}