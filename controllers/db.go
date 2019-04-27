package controllers

import (
	"github.com/bbhj/minabbs/models"
	"github.com/astaxie/beego"
)

type DBController struct {
	beego.Controller
}

// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /init [get]
func (u *DBController) InitDB() {
	beego.Info("===init db ==")

	var msg models.RetMsg
	flag, dbmsg :=  models.Init() 
	msg.Data = dbmsg
	if flag {
		msg.Status = 0
		msg.Msg = "init db successfully!"
	} else {
		msg.Status = -1
		msg.Errmsg = "init db failed!"
	}
	u.Data["json"] = msg
	u.ServeJSON()
}

// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /initdata [get]
func (u *DBController) InitData() {
	var reply models.Reply

	reply.Content = "回复测试Sed vel quisquam similique facilis fugit non."
	reply.TopicID = 5
	reply.UserID = 2
	models.AddReply(reply)

	var msg models.RetMsg
	msg.Status = 0
	msg.Msg = "init db successfully!"
	u.Data["json"] = msg
	u.ServeJSON()
}
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /test [get]
func (u *DBController) Test() {
	models.TestRelated()
	// u.Data["json"] = msg
	u.ServeJSON()
}
