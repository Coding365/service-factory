package test

import (
	"path/filepath"
	"runtime"
	"service-factory/k8s"
	_ "service-factory/routers"
	"testing"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	//k8s.TestCreateDeploy()
	//fmt.Println("fererererer")
	//k8s.TestJsonAndObj()
	k8s.TestK8sUtil()

}
