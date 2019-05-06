package controllers

import (
	"github.com/bbhj/minabbs/models"
	_ "encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Category
type CategoryController struct {
	beego.Controller
}

// @Title Create Category
// @Description create users
// @Param	body		body 	models.Category		true		"body for category content"
// @Success 200 {int} models.Category.Id
// @Failure 403 body is empty
// @router / [post]
func (u *CategoryController) Post() {
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Categories
// @Success 200 {object} models.Category
// @router / [get]
func (u *CategoryController) GetAll() {
	ulist, _ := models.GetAllCategories()
	var c models.Categories
	c.Data = ulist
	u.Data["json"] = c
	u.ServeJSON()
}
