package controllers

import (
	"gobook/models"
	"github.com/astaxie/beego"
	"gobook/filters"
	"gobook/utils"
)
type AdminController struct {
	beego.Controller
}

func (this *AdminController) URLMapping() {
	this.Mapping("/admin/index", this.Index)
}

// @router /admin/index [get]
func (this *AdminController) Index(){
	IsLogin, _ := filters.IsLogin(this.Ctx)
	if !IsLogin {
		this.Redirect("/login", 302)
		return
	}
	// 展示图书列表
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.BookPage(page.PageNo,page.PageSize)
	this.TplName = "admin/index.html"
}
