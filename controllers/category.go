package controllers

import (
	"github.com/bbhj/minabbs/models"
	_ "encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type CategoryController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *CategoryController) Post() {
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *CategoryController) GetAll() {
	ulist, _ := models.GetAllCategories()
	var c models.Categories
	c.Data = ulist
	u.Data["json"] = c
	u.ServeJSON()
}
