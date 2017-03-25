package test

import (
	"testing"
	"gobook/service"
	"github.com/astaxie/beego"
)

func Test_GetBookInfo(t *testing.T) {
	beego.Info(service.GetBookInfo("http://www.biquge.tw/81_81260/"))
}