package service

import (
	v1 "k8s.io/api/apps/v1"
	"service-factory/k8s"
	"service-factory/module/svc/dao"
	"service-factory/module/svc/entity"
)

type SvInfoService struct {
}

func NewSvInfoService() *SvInfoService {
	return new(SvInfoService)
}

func (this *SvInfoService) SvInfoToDeployment(svInfo *entity.SvInfo) *entity.Deployment {
	deployment := new(entity.Deployment)
	deployment.Port = 80
	deployment.Image = svInfo.Image
	deployment.Replicas = 1
	deployment.Code = svInfo.Code
	return deployment
}

func (this *SvInfoService) CreateService(svInfo *entity.SvInfo) *entity.Msg {
	srcDeployment := this.SvInfoToDeployment(svInfo)
	k8sutil := k8s.NewK8sUtil()
	clientset, err := k8sutil.GetK8sClient()
	if err != nil {
		panic(err)
		return entity.GetMsgSucc()
	}
	targetDeployment := v1.Deployment{}
	//生成deployment对象
	k8sutil.YamlToK8sResourceWithPath("module/svc/template/deployment.yaml", srcDeployment, &targetDeployment)
	_, err = clientset.AppsV1().Deployments("default").Create(&targetDeployment)
	if err != nil {
		panic(err)
		return entity.GetMsgSucc()
	}
	dao.SvInfoInsert(svInfo)
	return entity.GetMsgSucc()
}
