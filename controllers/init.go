package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bbhj/minabbs/models"
	"time"
)

type InitController struct {
	beego.Controller
}

// @Title Init Controller
// @Description Init tables
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RetMsg
// @Failure 403
// @router /db [get]
func (u *InitController) InitDB() {
	beego.Info("===init db ==")

	var msg models.RetMsg
	flag, dbmsg := models.Init()
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

// @Title Init data
// @Description Init base data
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /data [get]
func (u *InitController) InitData() {

	var c1 models.Category
	c1.Name = "公告"
	c1.Description = "公告信息"
	models.AddCategory(c1)

	var a1 models.Article
	a1.Title = "Hello, 欢迎使用本论坛！"
	a1.Body = "这是第一个帖子"
	a1.UserID = 1
	// a1.Category = 1
	models.AddArticle(a1)

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

// @Title Drop table
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /droptable [get]
func (u *InitController) Droptable() {
	models.Droptable()
	// u.Data["json"] = msg
	u.ServeJSON()
}

// @Title Uptime
// @Description App's uptime
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Babyinfo
// @Failure 403
// @router /run [get]
func (u *InitController) Run() {
	cost := time.Since(models.StartTime)
	beego.Error("======", cost)
	costs := fmt.Sprintf("%v", cost)
	// msg, _ := models.GetAllBabyinfo()
	// u.Data["json"] = msg
	// u.Data["json"] = map[string]time.Duration{"uptime": cost}
	u.Data["json"] = map[string]string{"uptime": costs}
	// u.Ctx.WriteString(costs)
	u.ServeJSON()
}
