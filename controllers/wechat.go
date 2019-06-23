package controllers

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/url"
	"time"

	_ "encoding/json"
	_ "reflect"

	_ "fmt"
	_ "strings"
	_ "time"

	"github.com/bbhj/bbac/models"
	localModel "github.com/bbhj/minabbs/models"
	"github.com/esap/wechat"
	"github.com/imroc/req"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	bm cache.Cache
)

func init() {
	bm, _ = cache.NewCache("memory", `{"interval":60}`)
}

//  /getphonenum  获取微信授权手机号

// Operations about WechatApi
type WechatController struct {
	beego.Controller
}

// @Title PcLoginState
// @Description get all Topics
// @Success 200 {object} models.Topic
// @router /pc-redirect-login-qr [get]
func (u *WechatController) PcLoginState() {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 64; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	bm.Put(string(result), nil, 300*time.Second)
	u.Data["json"] = map[string]interface{}{
		"data": string(result),
		"state": map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "",
		},
	}
	u.ServeJSON()
}

// @Title PcRedirectLoginQR
// @Description get all Topics
// @Success 200 {object} models.Topic
// @router /pc-redirect-login-qr [get]
func (u *WechatController) PcRedirectLoginQR() {
	beego.Info("=====", "WeChat PcRedirectLoginQR")
	redirect := u.Input().Get("redirect")
	state := u.Input().Get("state")

	params := url.Values{}
	params.Add("appid", "wxecc9f0214a302c3c")
	params.Add("redirect_uri", "http://127.0.0.1:8080/v1/api/wechat/pc-redirect?redirect="+redirect)
	params.Add("response_type", "code")
	params.Add("scope", "snsapi_base")
	params.Add("state", state)

	url := "https://open.weixin.qq.com/connect/oauth2/authorize?" + params.Encode() + "#wechat_redirect"
	u.Redirect(url, 302)
}

// @Title PcRedirect
// @Description get all Topics
// @Success 200 {object} models.Topic
// @router /pc-redirect [get]
func (u *WechatController) PcRedirect() {
	beego.Info("=====", "WeChat PcRedirect")
	code := u.Input().Get("code")
	state := u.Input().Get("state")

	params := url.Values{}
	params.Add("appid", "wxecc9f0214a302c3c")
	params.Add("secret", "14b1dea72482323578b322f9a042b3e2")
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	params.Add("state", state)

	r, err := req.Get("https://api.weixin.qq.com/sns/oauth2/access_token?" + params.Encode())
	if err != nil {
		beego.Error(err)
	}

	var result models.OpenWeixinAccessToken
	r.ToJSON(&result)
	if len(result.Openid) != 0 {
		err = bm.Put(state, result, 300*time.Second)
		beego.Info("=====", "WeChat PcRedirect", result, err)
		u.Data["json"] = map[string]interface{}{
			"data": result,
			"state": map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "",
			},
		}
		u.ServeJSON()
		return
	}

	// url := redirect + "?" + params.Encode()
	// u.Redirect(url, 302)
}

// @Title PcCheck
// @Description get all Topics
// @Success 200 {object} models.Topic
// @router /pc-redirect [get]
func (u *WechatController) PcCheck() {
	beego.Info("=====", "WeChat PcCheck")

	var request localModel.RequestWeChatQRLoginCheck
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &request); err != nil {
		return
	}

	result, ok := bm.Get(request.State).(models.OpenWeixinAccessToken)
	beego.Info("=====", "WeChat PcCheck", request.State, result, ok)
	if ok && len(result.Openid) > 0 {
		u.Data["json"] = map[string]interface{}{
			"data": map[string]interface{}{
				"ok":    true,
				"token": request.State,
				"name":  result.Openid,
				"roles": []string{"root", "admin"},
			},
			"state": map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "",
			},
		}
		// bm.Delete(state)
	} else {
		u.Data["json"] = map[string]interface{}{
			"data": map[string]interface{}{
				"ok": false,
			},
			"state": map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "",
			},
		}
	}

	u.ServeJSON()
	return
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
