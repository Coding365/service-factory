package main

import (
	"service-factory/module/svc/controllers"
	_ "service-factory/module/svc/routers"

	"github.com/astaxie/beego"

	_ "service-factory/init"

	_ "service-factory/module/svc/entity"
)

func init() {
	beego.Router("/svInfo", &controllers.SvInfoController{})
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
