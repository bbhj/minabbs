package controllers

import (
	"time"
	"encoding/json"
	"github.com/bbhj/minabbs/models"

	"github.com/astaxie/beego"
	"github.com/imroc/req"
)

// Operations about Users
type AuthorizationController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *AuthorizationController) Post() {

	var scene models.WechatLoginScene
	json.Unmarshal(u.Ctx.Input.RequestBody, &scene)
	// beego.Error("=!!!!!!=scene==", string(u.Ctx.Input.RequestBody))
	// beego.Error("=!!!!!!=scene==", scene)

	apiurl := beego.AppConfig.String("wechatcode") + scene.Code

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	var id models.WechatID
	a.ToJSON(&id)
	beego.Error("openidid === ", id)

	scene.Openid = id.Openid
	beego.Error("==!!!!!==", scene)

	b, _ := models.AddUserByWechat(id.Openid)
	beego.Error("======", b)
	// models.AddWechatLoginScene(scene)
	u.Data["json"] = b
	u.ServeJSON()
}