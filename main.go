package main

import (
	_ "github.com/lattecake/blog/routers"
	"github.com/astaxie/beego"
	"github.com/lattecake/blog/models"
	"github.com/lattecake/blog/helper"
	"github.com/astaxie/beego/logs"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var logFileName = beego.AppConfig.String("logs_path") + "blog.lattecake.com.log"
	beego.SetLogger(logs.AdapterFile, `{"filename": "`+logFileName+`","daily":true,"perm":"0644", "maxdays":365, "rotateperm": "0644"}`)
	logs.EnableFuncCallDepth(true)

	if beego.AppConfig.String("runmode") != "dev" {
		beego.BeeLogger.DelLogger("console")
		logs.Async()
	}
}

func main() {

	models.NewModel()

	beego.AddFuncMap("replace_image_url", helper.ReplaceImageUrl)
	beego.AddFuncMap("replace_image_src", helper.ReplaceImageSrc)
	beego.AddFuncMap("replace_post_image_src", helper.PostImageSrc)

	beego.Run()
}
