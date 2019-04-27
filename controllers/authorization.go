package controllers

import (
	"fmt"
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
	beego.Error("=!!!!!!=scene==", scene)

	apiurl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wxc4e11081e3d5bdf7&secret=e3284a3123d4ed4e06ee09bf0171bef7&js_code=%s&grant_type=authorization_code", scene.Code)

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	var id models.WechatID
	a.ToJSON(&id)
	beego.Error(id)

	scene.Openid = id.Openid
	beego.Error("==!!!!!==", scene)

	// models.AddWechatLoginScene(scene)
	u.Data["json"] = id
	u.ServeJSON()
}
