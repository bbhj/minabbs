package controllers

import (
	"github.com/bbhj/minabbs/models"
	"github.com/astaxie/beego"
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
	ret.State.Status= "success"
	ret.State.Message = ""

	// u.Data["json"] = map[string]string{"State": state}
	ret.Data = ulist
	u.Data["json"] = ret
	
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *LostController) Get() {

	uid, _ := u.GetInt(":uid")
	beego.Error("=====uid", uid)
	u.ServeJSON()
}
