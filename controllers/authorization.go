package controllers

import (
	"encoding/json"
	"github.com/bbhj/minabbs/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/imroc/req"
	"github.com/juusechec/jwt-beego"
)

// Operations about Authorization
type AuthorizationController struct {
	beego.Controller
}

// @Title Wechat User Login
// @Description Wechat Auth
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

	var profile models.Profile
	profile.UserID = b.ID
	profile.Openid = b.BoundWechat
	profile.AccessToken = "xxxxx"

	// models.AddWechatLoginScene(scene)
	u.Data["json"] = profile
	u.ServeJSON()
}

// @Title Wechat User auth
// @Description Wechat Auth
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /token [get]
func (u *AuthorizationController) Token() {
	et := jwtbeego.EasyToken{
		Username: "dean",
		Expires:  time.Now().Unix() + 3600, //Segundos
	}
	tokenString, _ := et.GetToken()

	et1 := jwtbeego.EasyToken{}
	valido, _, _ := et1.ValidateToken(tokenString)
	beego.Error("=======valido======", valido)
	if !valido {
		u.Ctx.Output.SetStatus(401)
		u.Data["json"] = "Permission Deny"
		u.ServeJSON()
	}

	u.Data["json"] = "{'tokenString': '" + tokenString + "'}"
	u.ServeJSON()
	return
}
