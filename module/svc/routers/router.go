// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"service-factory/module/svc/controllers"
)

func init() {
	beego.Router("/svInfo", &controllers.SvInfoController{})
	beego.Router("/svInfo/:id", &controllers.SvInfoController{})
	beego.Router("/svInfo/list", &controllers.SvInfoController{}, "get:GetList")
	beego.Router("/k8s/test", &controllers.K8sController{}, "get:TestK8s")
}
