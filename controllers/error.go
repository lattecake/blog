package controllers

import "github.com/astaxie/beego/logs"

// ErrorController operations for Error
type ErrorController struct {
	BaseController
}

// 系统报错 404
func (c *ErrorController) Error404() {
	logs.Notice(c.WorkerId(), " | page not found.")
	c.TplName = "error/404.tpl"
}

// 系统报错 501
func (c *ErrorController) Error501() {
	logs.Critical(c.WorkerId(), " | page status 501.")
	c.TplName = "error/50x.tpl"
}

// 系统报错 502
func (c *ErrorController) Error502() {
	logs.Critical(c.WorkerId(), " | page status 502.")
	c.TplName = "error/50x.tpl"
}

// 系统报错 数据库错误
func (c *ErrorController) ErrorDb() {
	logs.Critical(c.WorkerId(), " | db error.")
	c.TplName = "error/db.tpl"
}
