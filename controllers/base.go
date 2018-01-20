package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"github.com/astaxie/beego/logs"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
}

// 执行方法之前
func (c *BaseController) Prepare() {

	beginTime, _ := strconv.Atoi(c.Ctx.Input.Param("Begin-Time"))

	controllerName, actionName := c.GetControllerAndAction()

	c.Data["Active"] = controllerName + "_" + actionName
	c.Layout = "default_layout.tpl"
	c.LayoutSections = make(map[string]string)

	prepareTime := time.Now().UnixNano() - int64(beginTime)

	logs.Info("| Worker-Id | ", c.WorkerId(), " | ", controllerName+"_"+actionName)
	logs.Info("| Worker-Id | ", c.Ctx.Input.Param("Worker-Id"), " | Begin-Time | ", time.Unix(0, int64(beginTime)).Format("2006-01-02 15:04:05.999999"), " | Prepare | ", time.Duration(prepareTime))
}

// 执行完方法之后
func (c *BaseController) Finish() {

	beginTime := c.Ctx.Input.Param("Begin-Time")
	begin, _ := strconv.Atoi(beginTime)
	end, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano(), 10))

	//c.Ctx.Output.Status = 200
	c.Ctx.Output.Header("x-response-id", c.Ctx.Request.Header.Get("Worker-Id"))
	c.Ctx.Output.Header("x-response-time", time.Duration(end - begin).String())

	logs.Info("| Worker-Id | ", c.Ctx.Input.Param("Worker-Id"), " | status | 200 | Total-Time | ", time.Duration(end-begin))
}

func (c *BaseController) WorkerId() string {
	return c.Ctx.Input.Param("Worker-Id")
}
