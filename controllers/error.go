package controllers

import "github.com/astaxie/beego/logs"

// ErrorController operations for Error
type ErrorController struct {
	BaseController
}

// 系统报错 404
func (c *ErrorController) Error404() {
	c.Data["errorCode"] = 404
	c.Data["message"] = "我迷失了,找不到前进的方向"
	logs.Notice(c.WorkerId(), " | page not found.")
	c.TplName = "error/404.tpl"
}

// 系统报错 501
func (c *ErrorController) Error501() {
	c.Data["errorCode"] = 501
	c.Data["message"] = "系统失恋崩溃了"
	logs.Critical(c.WorkerId(), " | page status 501.")
	c.TplName = "error/50x.tpl"
}

// 系统报错 502
func (c *ErrorController) Error502() {
	c.Data["errorCode"] = 502
	c.Data["message"] = "系统失恋崩溃了"
	logs.Critical(c.WorkerId(), " | page status 502.")
	c.TplName = "error/50x.tpl"
}

// 系统报错 503
func (c *ErrorController) Error503() {
	logs.Critical(c.WorkerId(), " | page status 502.")
	c.Data["errorCode"] = 503
	c.Data["message"] = "系统失恋崩溃了"
	c.TplName = "error/50x.tpl"
}

// 系统报错 500
func (c *ErrorController) Error500() {
	c.Data["errorCode"] = 500
	c.Data["message"] = "系统失恋崩溃了"
	logs.Critical(c.WorkerId(), " | page status 502.")
	c.TplName = "error/50x.tpl"
}

// 系统报错 数据库错误
func (c *ErrorController) ErrorDb() {
	c.Data["errorCode"] = 500
	c.Data["message"] = "数据库跑路了"
	logs.Critical(c.WorkerId(), " | db error.")
	c.TplName = "error/db.tpl"
}
