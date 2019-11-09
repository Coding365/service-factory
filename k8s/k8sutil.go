package k8s

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ghodss/yaml"
	"html/template"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

type K8sUtil struct {
}

type K8sUtilInterface interface {
	GetK8sClient() (*kubernetes.Clientset, error)
}

func NewK8sUtil() *K8sUtil {
	return new(K8sUtil)
}

func (*K8sUtil) GetK8sClient() (*kubernetes.Clientset, error) {

	k8shost := beego.AppConfig.String("k8shost")

	config, err := clientcmd.BuildConfigFromFlags(k8shost, "")
	if err != nil {
		logs.Error("")
		panic(err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset, err
}

func (this *K8sUtil) YamlToK8sResourceWithPath(filePath string, source interface{}, target interface{}) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
		return err
	}
	yamlByte, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
		return err
	}

	t := template.New("deployment")

	t, _ = t.Parse(string(yamlByte))
	buf := new(bytes.Buffer)

	t.Execute(buf, &source)

	return this.YamlToK8sResource(buf.Bytes(), target)
}

func (*K8sUtil) YamlToK8sResource(data []byte, res interface{}) error {
	jsonByte, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, &res)
	if err != nil {
		return err
	}

	return nil
}
