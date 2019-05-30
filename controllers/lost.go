package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bbhj/minabbs/models"
)

// Operations about Users
type LostController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *LostController) GetAll() {
	ulist, _ := models.GetAllBabyinfo()
	u.Data["json"] = ulist
	// var state models.State
	var ret models.RetDataList
	ret.State.Code = 200
	ret.State.Status = "success"
	ret.State.Message = ""

	// u.Data["json"] = map[string]string{"State": state}
	ret.Data = ulist
	u.Data["json"] = ret

	u.ServeJSON()
}

// @Title Get
// @Description get lost info by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /detail/:id [get]
func (u *LostController) Get() {

	id, _ := u.GetInt(":id")
	info, _ := models.GetBabyinfoById(id)

	var ret models.RetDataList
	ret.State.Code = 200
	ret.State.Status = "success"
	ret.State.Message = ""

	ret.Data = info
	u.Data["json"] = ret
	u.ServeJSON()
}