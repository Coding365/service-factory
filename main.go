package main

import (
	"service-factory/controllers"
	_ "service-factory/routers"

	"github.com/astaxie/beego"

	_ "service-factory/init"

	_ "service-factory/models"
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
