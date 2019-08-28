package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"service-factory/models"
)

// Operations about object
type SvInfoController struct {
	beego.Controller
}

func (s *SvInfoController) Post() {
	var svInfo models.SvInfo
	json.Unmarshal(s.Ctx.Input.RequestBody, &svInfo)
	uid := models.Insert(&svInfo)
	s.Data["json"] = map[string]string{"uid": uid}
	s.ServeJSON()
}

func (s *SvInfoController) Get() {
	id := s.Ctx.Input.Query("id")
	svInfo := models.Select(id)
	s.Data["json"] = svInfo
	s.ServeJSON()
}
