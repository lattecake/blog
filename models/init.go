package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	var host = beego.AppConfig.String("mysql_host")
	var port = beego.AppConfig.String("mysql_port")
	var user = beego.AppConfig.String("mysql_user")
	var password = beego.AppConfig.String("mysql_password")
	var database = beego.AppConfig.String("mysql_db")

	logs.Info("init db pool")

	orm.RegisterDataBase("default", "mysql", ""+user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&loc=Asia%2FShanghai")

}

func NewModel() {
	orm.Debug = true // TODO DEBUG
}
