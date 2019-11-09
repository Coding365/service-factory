package controllers

import (
	"github.com/astaxie/beego"
	"service-factory/k8s"
)

type K8sController struct {
	beego.Controller
}

func (this *K8sController) TestK8s() {
	k8s.TestCreateDeploy()
}
