package routers

import (
	"gobook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.LinkController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.BookController{})
	beego.Include(&controllers.ChapterController{})
	beego.Include(&controllers.ChapterLogController{})
	beego.Include(&controllers.QrCodeController{})
	beego.Include(&controllers.JobController{})
	beego.Include(&controllers.TodayOnhistory{})
	beego.Include(&controllers.WxAppController{})
	beego.Include(&controllers.ShellJobController{})
}
