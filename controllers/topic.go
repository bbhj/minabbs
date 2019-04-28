package controllers

import (
	"encoding/json"
	"github.com/bbhj/minabbs/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type TopicController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TopicController) Post() {
	var article models.Article
	json.Unmarshal(u.Ctx.Input.RequestBody, &article)
	// remove(article, article.User)
	beego.Error("======", article)
	a, _ := models.AddArticle(article)
	beego.Error("======", a)
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *TopicController) GetAll() {
	page := u.GetString("page")
	include := u.GetString("include")
	beego.Info("=====", page, include)

	var topics models.Topics
	var art2 models.Topic

	articles, _ := models.GetAllArticles()
	for _, article := range articles {
		art2.Article = article
		art2.Category, _ = models.GetCategory(article.CategoryID)
		art2.User, _ = models.GetUser(article.UserID)
		// // art2.TopReplies.Data.Reply, _ = models.GetAllReplies()
		// replies, _  :=  models.GetAllReplies()
		// var ru models.ReplyUser
		// for _, reply := range replies {
		// 	ru.Reply  = reply
		// 	ru.User, _ = models.GetUser(reply.UserID)
		// 	art2.TopReplies.Data = append(art2.TopReplies.Data, ru)
		// }
		topics.Data = append(topics.Data, art2)

	}
	// beego.Info("=====", topics.Data)

	u.Data["json"] = topics
	u.ServeJSON()
}

// @Title Get
// @Description get user by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (u *TopicController) Get() {

	topic_id, _ := u.GetInt(":id")
	// Article2
	var topic models.Topic
	// json.Unmarshal([]byte(aa), &topic)

	topic.Article, _ = models.GetArticle(topic_id)

	topic.Category, _ = models.GetCategory(3)
	// topic.TopReplies.Data, _ = models.GetAllReplies()
	topic.User, _ = models.GetUser(topic.Article.UserID)

	//	replies, _  :=  models.GetAllReplies()
	replies, _ := models.GetRepliesByTopicID(topic_id)
	var ru models.ReplyUser
	for _, reply := range replies {
		ru.Reply = reply
		ru.User, _ = models.GetUser(reply.UserID)
		topic.TopReplies.Data = append(topic.TopReplies.Data, ru)
	}

	// for i, _ := range topic.TopReplies.Data {
	// 	beego.Error("index=====", i)
	// 	topic.TopReplies.Data[i].User, _ = models.GetUser(topic.TopReplies.Data[i].UserID)
	// }

	u.Data["json"] = topic
	u.ServeJSON()
}

// @Title 帖子回复
// @Description get user by uid
// @Param       uid             path    string  true            "The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid/replies [Post]
func (u *TopicController) PostReply() {
	var reply models.Reply
	json.Unmarshal(u.Ctx.Input.RequestBody, &reply)
	beego.Error(reply.TopicID)
	models.AddReply(reply)
	u.ServeJSON()
}

// @Title Get
// @Description get topic replies  by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:id/replies [get]
func (u *TopicController) GetReply() {
	topic_id, _ := u.GetInt(":id")
	var rep models.TopicReplies
	// rep.replies, _ = models.GetAllReplies()
	// rep.Data, _ = models.GetAllReplies()
	replies, _ := models.GetRepliesByTopicID(topic_id)

	var replyuser models.ReplyUser
	for _, reply := range replies {
		replyuser.Reply = reply
		replyuser.User, _ = models.GetUser(reply.UserID)
		rep.Data = append(rep.Data, replyuser)
	}
	// rep.Data[0].User, _ = models.GetUser(1)
	// rep.Data[1].User, _ = models.GetUser(1)
	rep.Meta.Pagination.Count = 2
	rep.Meta.Pagination.Total = 2
	rep.Meta.Pagination.Links = ""
	rep.Meta.Pagination.TotalPages = 2
	u.Data["json"] = rep
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *TopicController) Put() {
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *TopicController) Delete() {
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *TopicController) Login() {
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *TopicController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
