package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/book/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "SaveOrUpdate",
			Router: `/book/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "TaskUpdate",
			Router: `/book/taskUpdate/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/book/delete/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "UrlInfo",
			Router: `/book/url/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/book/search`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/book/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "ListJson",
			Router: `/book/json/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "Export",
			Router: `/book/export/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:BookController"] = append(beego.GlobalControllerRouter["gobook/controllers:BookController"],
		beego.ControllerComments{
			Method: "LocalUpdate",
			Router: `/book/localUpdate/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/chapter/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/chapter/detail/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "DetailJson",
			Router: `/chapter/json/detail/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Next",
			Router: `/chapter/next/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Pre",
			Router: `/chapter/pre/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "NextJson",
			Router: `/chapter/json/next/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "PreJson",
			Router: `/chapter/json/pre/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "SaveOrUpdate",
			Router: `/chapter/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/chapter/delete/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "DeleteBook",
			Router: `/chapter/deletebook/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/chapter/list/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "ListJson",
			Router: `/chapter/json/list/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "HasNewChapter",
			Router: `/chapter/new/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "ListByLog",
			Router: `/chapter/list/:tag(\w+)/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "FindByTitle",
			Router: `/chapter/title/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/chapter/update/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterLogController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterLogController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/log/chapter/:tag([\w]+)/:bookId([0-9]+)/:index([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ChapterLogController"] = append(beego.GlobalControllerRouter["gobook/controllers:ChapterLogController"],
		beego.ControllerComments{
			Method: "FindTag",
			Router: `/log/tag/:tag([\w]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/job/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "SaveOrUpdate",
			Router: `/job/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/job/delete/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/job/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "Start",
			Router: `/job/start/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "Pause",
			Router: `/job/pause/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:JobController"] = append(beego.GlobalControllerRouter["gobook/controllers:JobController"],
		beego.ControllerComments{
			Method: "Run",
			Router: `/job/run/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/link/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/link/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/link/delete/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/link/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/link/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "PostLink",
			Router: `/link/ajax/post`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:LinkController"] = append(beego.GlobalControllerRouter["gobook/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Import",
			Router: `/link/ajax/import`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"] = append(beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/proxyIp/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"] = append(beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/proxyIp/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"] = append(beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/proxyIp/delete/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"] = append(beego.GlobalControllerRouter["gobook/controllers:ProxyIpController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/proxyIp/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:QrCodeController"] = append(beego.GlobalControllerRouter["gobook/controllers:QrCodeController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/qr/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:QrCodeController"] = append(beego.GlobalControllerRouter["gobook/controllers:QrCodeController"],
		beego.ControllerComments{
			Method: "Decode",
			Router: `/qr/decode`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:ShellJobController"] = append(beego.GlobalControllerRouter["gobook/controllers:ShellJobController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/shell/job/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:TodayOnhistory"] = append(beego.GlobalControllerRouter["gobook/controllers:TodayOnhistory"],
		beego.ControllerComments{
			Method: "QueryEvent",
			Router: `/todayOnhistory/queryEvent`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:TodayOnhistory"] = append(beego.GlobalControllerRouter["gobook/controllers:TodayOnhistory"],
		beego.ControllerComments{
			Method: "QueryDetail",
			Router: `/todayOnhistory/queryDetail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
		beego.ControllerComments{
			Method: "LoginPage",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
		beego.ControllerComments{
			Method: "RegisterPage",
			Router: `/register`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/user/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:WxAppController"] = append(beego.GlobalControllerRouter["gobook/controllers:WxAppController"],
		beego.ControllerComments{
			Method: "UserInfo",
			Router: `/wxapp/userInfo`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gobook/controllers:WxAppController"] = append(beego.GlobalControllerRouter["gobook/controllers:WxAppController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/wxapp/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
