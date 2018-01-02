package controllers

import (
	"github.com/astaxie/beego"
	"gobook/models"
	"gobook/utils"
	"strconv"
	"gobook/service"
	"gobook/cache"
)

type BookController struct {
	beego.Controller
}
//bm := NewBeeMap()
//go get github.com/astaxie/beego/cache

func (this *BookController) URLMapping() {
	this.Mapping("/book/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/book/save", this.SaveOrUpdate)
	this.Mapping("/book/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/book/list", this.List)
	this.Mapping("/book/json/list", this.ListJson)
	this.Mapping("/book/taskUpdate/:id([0-9]{0,})", this.TaskUpdate)
	this.Mapping("/book/url/info", this.UrlInfo)
	this.Mapping("/book/search", this.Search)
	this.Mapping("/book/export", this.Export)
	this.Mapping("/book/localUpdate", this.LocalUpdate)
}

// @router /book/edit/:id([0-9]{0,}) [get]
func (this *BookController) Edit() {
	id:=this.Ctx.Input.Param(":id")
	bk:= models.Book{}
	if id!=""{
		if i,err:=strconv.ParseInt(id, 10, 64); err==nil{
			ok,book:=models.FindBookById(i)
			if ok{
				bk=book
			}
		}
	}
	this.Data["entry"] = bk
	this.TplName = "book/edit.tpl"
}

// @router /book/save [post]
func (this *BookController) SaveOrUpdate() {
	book := models.Book{}
	if err := this.ParseForm(&book); err != nil {
		beego.Error(err)
	}
	models.BookSaveOrUpdate(&book)
	this.Redirect("/book/list", 302)
}
// @router /book/taskUpdate/:id([0-9]{0,}) [get]
func (this *BookController) TaskUpdate() {
	id:=this.Ctx.Input.Param(":id")
	json:=JsonObj{Code: -1, Msg:"error"}
	if id!=""{
		if i,err:=strconv.ParseInt(id, 10, 64); err==nil{
			ok,book:=models.FindBookById(i)
			if ok{
				tag:="cache_book_"+id
				if isRunning:=service.UpdateBook(&book,tag);isRunning{
					json = JsonObj{Code: -1, Msg:tag+"=正在更新！更新时间="+cache.Get(tag).(string)}
				}else{
					json = JsonObj{Code: 0, Msg:"ok"}
				}
			}
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}

// @router /book/delete/:id([0-9]+)
func (this *BookController) Delete(){
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		if i,err:=strconv.Atoi(id); err==nil{
			book := &models.Book{Id : i}
			models.BookDelete(book)
		}

	}

	jsonMap := map[string]interface{}{
		"code":0,
		"msg":"ok",
	}
	this.Data["json"] = jsonMap
	this.ServeJSON()
}
// @router /book/url/info
func (this *BookController) UrlInfo(){
	url:=this.GetString("Url")
	book:=&models.Book{}
	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"no result",
		"result":nil,
	}
	if url!=""{
		book=service.GetBookInfo(url)
		jsonMap=map[string]interface{}{
			"code":0,
			"msg":"",
			"result":book,
		}
	}

	this.Data["json"] = jsonMap
	this.ServeJSON()
}

// @router /book/search
func (this *BookController) Search(){
	name:=this.GetString("Name")
	book:=&models.Book{}
	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"no result",
		"result":nil,
	}
	if name!=""{
		book=service.Search(name)
		jsonMap=map[string]interface{}{
			"code":0,
			"msg":"",
			"result":book,
		}
	}

	this.Data["json"] = jsonMap
	this.ServeJSON()
}

// @router /book/list
func (this *BookController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["runmode"] = beego.AppConfig.String("runmode")
	this.Data["page"] = models.BookPage(page.PageNo,page.PageSize)
	this.TplName = "book/list.tpl"
}
// @router /book/json/list
func (this *BookController) ListJson() {
	jsonMap:=map[string]interface{}{
		"code":0,
		"msg":"",
	}

	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}

	jsonMap["page"] = models.BookPage(page.PageNo,page.PageSize)
	this.Data["json"] = jsonMap
	this.ServeJSON()
}

// @router /book/export/:id([0-9]+)
func (this *BookController) Export() {

	id:=this.Ctx.Input.Param(":id")
	err:=models.Export(id)
	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"",
	}
	if err!=nil{
		jsonMap["msg"]=err.Error()
	}else{
		jsonMap["code"]=0
		jsonMap["msg"]="ok"
	}

	this.Data["json"] = jsonMap
	this.ServeJSON()
}
// @router /book/localUpdate/:id([0-9]+)
func (this *BookController) LocalUpdate() {

	id:=this.Ctx.Input.Param(":id")
	err,chapters,rule:=models.FindEmptyChapters(id)

	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"",
	}
	if err!=nil{
		jsonMap["msg"]=err.Error()
	}else{
		go (func() {

			service.GetChapterContent(rule,chapters,100)

			/*for _,chapter:=range chapters{
				beego.Info("抓取="+chapter.Title)
				content:=service.GetContent(chapter.Url,rule)
				chapter.Content=content
			}*/
			beego.Info("<<<<<<<<<<<<<<<<<")
			models.UpdateChapterContent(chapters)
			beego.Info(">>>>>>>>>>>>>>>>>")
		})()
		jsonMap["code"]=0
		jsonMap["msg"]="ok"
	}

	this.Data["json"] = jsonMap
	this.ServeJSON()
}
