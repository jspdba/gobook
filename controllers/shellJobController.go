package controllers

import (
	"github.com/astaxie/beego"
	"gobook/service"
)

type ShellJobController struct {
	beego.Controller
}

func (this *ShellJobController) URLMapping() {
	this.Mapping("/shell/job/list", this.List)
}
// @router /shell/job/list
func(this *ShellJobController) List() {
	jobs := service.ListShellJobs()
	this.Data["jobs"] = jobs
	this.TplName = "shellJob/list.tpl"
}
