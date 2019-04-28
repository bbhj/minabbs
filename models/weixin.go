package models

import (
	_ "errors"
	"github.com/jinzhu/gorm"
)

var ()

type (
	WechatLoginScene struct {
		gorm.Model
		Code   string `json:code"`
		Openid string `json:"openid"`
		Scene  int    `json:"scene"`
		Query 		`json:"query"`
		Path   string `json:"path"`
		IV            string `json:"iv"`
		EncryptedData string `json:"encryptedData"`
		User
	}

	Query struct {
		ID string `json:"id"`
		Openid string `json:"openid"`
	}

	WechatID struct {
		gorm.Model
		Openid        string `json:"openid"`
		Unionid       string `json:"unionid"`
		AccessToken   string `json:"access_token"`
		ExpiresIn     int    `json:"expires_in"`
		RefreshToken  string `json:"refresh_token"`
		Scope         string `json:"scope"`
		SessionKey    string `json:"session_key"` // 小程序登录session_key
	}
)