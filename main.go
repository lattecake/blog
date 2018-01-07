package main

import (
	_ "github.com/lattecake/blog/routers"
	"github.com/astaxie/beego"
	"github.com/lattecake/blog/models"
	"github.com/lattecake/blog/helper"
)

func main() {

	models.NewModel()

	beego.AddFuncMap("replace_image_url", helper.ReplaceImageUrl)
	beego.AddFuncMap("replace_image_src", helper.ReplaceImageSrc)
	beego.AddFuncMap("replace_post_image_src", helper.PostImageSrc)

	beego.Run()
}
