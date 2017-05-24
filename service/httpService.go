package service

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"github.com/astaxie/beego"
)

func Get(upath string,v *url.Values) (content []byte,err error){
	u, _ := url.Parse(upath)
	//q := u.Query()
	//q.Set("username", "user")
	//q.Set("password", "passwd")
	u.RawQuery = v.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		beego.Error(err)
		return content,err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		beego.Error(err)
		return content,err
	}
	beego.Info(string(result))
	return result,err
}
