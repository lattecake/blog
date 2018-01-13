package routers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
	"github.com/satori/go.uuid"
	"github.com/lattecake/utils"
	"github.com/lattecake/blog/controllers"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		workerId := uuid.NewV4().String()
		//u, _ := uuid.NewV4()
		//workerId := u.String()

		beginTime := strconv.FormatInt(time.Now().UnixNano(), 10)

		logs.Info(" | ", ctx.Request.RemoteAddr, " | ", ctx.Input.IP(), " | Request-Id | ", ctx.Request.Header.Get("x-request-id"), " | Source-Time | ", ctx.Request.Header.Get("Source-Time"), "| Worker-Id | ", workerId, " | Server-IP | ", utils.LocalAddress(), " | Method | ", ctx.Request.Method, " | URI | ", ctx.Request.RequestURI)
		ctx.Input.SetParam("Worker-Id", workerId)
		ctx.Input.SetParam("Begin-Time", beginTime)
		//ctx.Input.SetParam("count", "0")
	})

	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.PostController{})
	beego.Include(&controllers.HomeController{})
	//beego.Router("/post/:pageId", &controllers.PostController{}, "get:GetAll")
	beego.Router("/learn/:pageId", &controllers.PostController{}, "get:GetAll")
	beego.Router("/posts", &controllers.PostController{}, "get:GetAll")
	beego.Router("/posts/:pageId", &controllers.PostController{}, "get:GetAll")
	beego.Router("/life/:pageId", &controllers.PostController{}, "get:GetAll")
	beego.Router("/life", &controllers.PostController{}, "get:GetAll")
	beego.SetStaticPath("/favicon.ico", "static")
}
