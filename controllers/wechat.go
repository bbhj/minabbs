package controllers

import (
	"encoding/base64"
	"encoding/json"
	_ "encoding/json"
	_ "reflect"

	_ "fmt"
	"github.com/bbhj/bbac/models"
	"github.com/esap/wechat"
	"github.com/imroc/req"
	_ "strings"
	_ "time"

	"github.com/astaxie/beego"
)

//  /getphonenum  获取微信授权手机号

// Operations about WechatApi
type WechatController struct {
	beego.Controller
}

// @Title Create Wechat QR Code
// @Description create QR Code
// @Param	body		body 	models.Article	true		"body for user content"
// @Success 200 {int} models.Article.Id
// @Failure 403 body is empty
// @router /createqrcode [post]
func (u *WechatController) CreateQRcode() {
	beego.Info("create qr code ====")
	var params models.QRCodeRequestParms
	var qrret models.QRCodeReturn

	json.Unmarshal(u.Ctx.Input.RequestBody, &params)
	beego.Error(params)

	apiurl := "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=" + wechat.GetAccessToken()

	param := req.Param{
		// "access_token": wechat.GetAccessToken(),
		"page":       params.Page,
		"scene":      "forward?user=1111",
		"width":      430,
		"auto_color": false,
		"is_hyaline": false,
	}

	// req.Debug = false

	// r, err := req.Post(apiurl, "application/json", req.BodyJSON(param))
	r, err := req.Post(apiurl, req.BodyJSON(param))
	if err != nil {
		beego.Error(err)
	}

	var ret models.RetMsg
	r.ToJSON(&qrret)
	if qrret.Errcode != 0 {
		beego.Error("get qr code from wechat failed, ", ret)
		ret.Status = -1
		ret.Errcode = qrret.Errcode
		ret.Errmsg = qrret.Errmsg

		u.Data["json"] = ret
		u.ServeJSON()
		return
	}
	// beego.Error(r)
	// Read entire JPG into byte slice.
	// reader := bufio.NewReader(r)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(r.Bytes())

	// beego.Error(encoded)

	// r.ToFile("qq3.jpg")

	u.Ctx.WriteString(encoded)
	return
}
