package main

import (
	_ "kyronApi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:nohoraluz@127.0.0.1:5433/hojas_de_vida?sslmode=disable&search_path=public")
}

func main() {
	orm.Debug = true
	if beego.BConfig.RunMode == "dev" {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	AllowHeaders: []string{"Origin", "x-requested-with",
	"content-type",
	"accept",
	"origin",
	"authorization",
	"x-csrftoken"},
	ExposeHeaders: []string{"Content-Length"},
	AllowCredentials: true,
	}))
	beego.Run()
}

