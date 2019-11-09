package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"service-factory/module/svc/dao"
	"service-factory/module/svc/entity"
	"service-factory/module/svc/service"
	"strconv"
)

// Operations about object
type SvInfoController struct {
	beego.Controller
}

func (this *SvInfoController) Post() {
	var svInfo entity.SvInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &svInfo)
	//uid := dao.SvInfoInsert(&svInfo)
	//this.Data["json"] = map[string]string{"uid": uid}
	svinfoservice := service.NewSvInfoService()
	msg := svinfoservice.CreateService(&svInfo)
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *SvInfoController) Get() {
	id := this.Ctx.Input.Param(":id")
	svInfo := dao.SvInfoSelect(id)
	this.Data["json"] = svInfo
	this.ServeJSON()
}

func (this *SvInfoController) Delete() {
	id := this.Ctx.Input.Param(":id")
	svInfo := dao.SvInfoSelect(id)
	this.Data["json"] = svInfo
	this.ServeJSON()

}

func (this *SvInfoController) GetList() {
	pageNoStr := this.Ctx.Input.Query("pageNo")
	pageSizeStr := this.Ctx.Input.Query("pageSize")

	pageNo, err1 := strconv.Atoi(pageNoStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil {
		pageNo = 1
		pageSize = 20
	}

	svInfos := dao.SvInfoSelectAll(pageNo, pageSize)
	msg := entity.GetMsgSucc()
	msg.SetMsgResult(svInfos)
	this.Data["json"] = msg
	this.ServeJSON()

}
