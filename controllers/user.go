package controllers

import (
	"fmt"
	"github.com/bbhj/minabbs/models"
	"time"
	"github.com/imroc/req"	
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var scene models.WechatLoginScene
	json.Unmarshal(u.Ctx.Input.RequestBody, &scene)
	beego.Error("=!!!!!!=scene==", scene.User)

	apiurl := beego.AppConfig.String("wechatcode") + scene.Code

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	var id models.WechatID

	a.ToJSON(&id)
	beego.Error("openidid === ", id)
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	ulist, _ := models.GetAllUsers()
	u.Data["json"] = ulist
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {

	uid, _ := u.GetInt(":uid")
	beego.Error("=====uid", uid)
	user, err := models.GetUser(uid)
	if err != nil {
		beego.Error("query user failed")
	}

	u.Data["json"] = user

	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param       uid             path    string  true            "The key for
// staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid/replies [get]
func (u *UserController) GetUserReply() {

	aa := `{"data":[{"id":65,"title":"Inventore ex voluptates delectus et.","body":"Nihil rem quia eum et quia. Qui tenetur culpa quas consectetur dignissimos velit dolorem. Id ut pariatur sint rerum. Odio autem ut tempora.","user_id":5,"category_id":3,"reply_count":0,"view_count":0,"last_reply_user_id":0,"excerpt":"Inventore ex voluptates delectus et.","slug":null,"created_at":"2019-03-03 21:40:48","updated_at":"2019-04-24 21:07:15","user":{"id":5,"name":"Peter McDermott V","email":"chad.armstrong@example.com","avatar":"https:\/\/iocaffcdn.phphub.org\/uploads\/images\/201710\/14\/1\/xAuDMxteQy.png?imageView2\/1\/w\/200\/h\/200","introduction":"Ex vel assumenda neque impedit nam.","bound_phone":false,"bound_wechat":false,"last_actived_at":"2019-03-13 15:05:17","created_at":"2019-03-13 15:05:17","updated_at":"2019-04-24 21:06:39"},"category":{"id":3,"name":"\u95ee\u7b54","description":"\u8bf7\u4fdd\u6301\u53cb\u5584\uff0c\u4e92\u5e2e\u4e92\u52a9"}},{"id":19,"title":"Quo dolore a doloremque consequatur quia.","body":"Aut ipsa veniam voluptatem nihil aut ut. Minima ut est eos qui. Hic illum accusamus nam possimus minima corrupti.  Deserunt voluptatem exercitationem et aut placeat minus maxime minima.","user_id":5,"category_id":4,"reply_count":0,"view_count":0,"last_reply_user_id":0,"excerpt":"Quo dolore a doloremque consequatur quia.","slug":null,"created_at":"2019-02-23 14:57:22","updated_at":"2019-02-23 17:02:54","user":{"id":5,"name":"Peter McDermott V","email":"chad.armstrong@example.com","avatar":"https:\/\/iocaffcdn.phphub.org\/uploads\/images\/201710\/14\/1\/xAuDMxteQy.png?imageView2\/1\/w\/200\/h\/200","introduction":"Ex vel assumenda neque impedit nam.","bound_phone":false,"bound_wechat":false,"last_actived_at":"2019-03-13 15:05:17","created_at":"2019-03-13 15:05:17","updated_at":"2019-04-24 21:06:39"},"category":{"id":4,"name":"\u516c\u544a","description":"\u7ad9\u70b9\u516c\u544a"}},{"id":69,"title":"Minima reiciendis quaerat et reprehenderit voluptas.","body":"Eveniet repudiandae ullam nemo cum. Dolores neque est quod quasi quis. Id eligendi eveniet sint nulla.","user_id":5,"category_id":3,"reply_count":0,"view_count":0,"last_reply_user_id":0,"excerpt":"Minima reiciendis quaerat et reprehenderit voluptas.","slug":null,"created_at":"2019-02-21 02:01:45","updated_at":"2019-02-27 01:45:05","user":{"id":5,"name":"Peter McDermott V","email":"chad.armstrong@example.com","avatar":"https:\/\/iocaffcdn.phphub.org\/uploads\/images\/201710\/14\/1\/xAuDMxteQy.png?imageView2\/1\/w\/200\/h\/200","introduction":"Ex vel assumenda neque impedit nam.","bound_phone":false,"bound_wechat":false,"last_actived_at":"2019-03-13 15:05:17","created_at":"2019-03-13 15:05:17","updated_at":"2019-04-24 21:06:39"},"category":{"id":3,"name":"\u95ee\u7b54","description":"\u8bf7\u4fdd\u6301\u53cb\u5584\uff0c\u4e92\u5e2e\u4e92\u52a9"}},{"id":68,"title":"Sapiente qui rem ea perspiciatis qui qui.","body":"Molestiae mollitia dolor sed molestiae in. Quis iure ad quis reprehenderit id quo doloremque. Ducimus corrupti omnis maxime veniam nemo assumenda autem.","user_id":5,"category_id":3,"reply_count":0,"view_count":0,"last_reply_user_id":0,"excerpt":"Sapiente qui rem ea perspiciatis qui qui.","slug":null,"created_at":"2019-02-16 15:33:36","updated_at":"2019-02-16 21:49:36","user":{"id":5,"name":"Peter McDermott V","email":"chad.armstrong@example.com","avatar":"https:\/\/iocaffcdn.phphub.org\/uploads\/images\/201710\/14\/1\/xAuDMxteQy.png?imageView2\/1\/w\/200\/h\/200","introduction":"Ex vel assumenda neque impedit nam.","bound_phone":false,"bound_wechat":false,"last_actived_at":"2019-03-13 15:05:17","created_at":"2019-03-13 15:05:17","updated_at":"2019-04-24 21:06:39"},"category":{"id":3,"name":"\u95ee\u7b54","description":"\u8bf7\u4fdd\u6301\u53cb\u5584\uff0c\u4e92\u5e2e\u4e92\u52a9"}},{"id":56,"title":"Dolorem occaecati in non perferendis nostrum qui.","body":"Quia mollitia doloribus deserunt dolorem. Magnam sint sed cum illum blanditiis. Dolore tenetur ut beatae aliquid delectus.","user_id":5,"category_id":2,"reply_count":0,"view_count":0,"last_reply_user_id":0,"excerpt":"Dolorem occaecati in non perferendis nostrum qui.","slug":null,"created_at":"2019-02-15 10:18:04","updated_at":"2019-02-16 12:06:12","user":{"id":5,"name":"Peter McDermott V","email":"chad.armstrong@example.com","avatar":"https:\/\/iocaffcdn.phphub.org\/uploads\/images\/201710\/14\/1\/xAuDMxteQy.png?imageView2\/1\/w\/200\/h\/200","introduction":"Ex vel assumenda neque impedit nam.","bound_phone":false,"bound_wechat":false,"last_actived_at":"2019-03-13 15:05:17","created_at":"2019-03-13 15:05:17","updated_at":"2019-04-24 21:06:39"},"category":{"id":2,"name":"\u6559\u7a0b","description":"\u5f00\u53d1\u6280\u5de7\u3001\u63a8\u8350\u6269\u5c55\u5305\u7b49"}},{"id":8,"title":"Rerum ut dolores vero iste facere et.","body":"Totam et voluptate enim ut accusamus laborum. Amet rerum id omnis molestiae sit pariatur.","user_id":5,"category_id":2,"reply_count":0,"view_count":0,"last_reply_user_id":0,"excerpt":"Rerum ut dolores vero iste facere et.","slug":null,"created_at":"2019-02-14 16:47:17","updated_at":"2019-02-24 00:12:40","user":{"id":5,"name":"Peter McDermott V","email":"chad.armstrong@example.com","avatar":"https:\/\/iocaffcdn.phphub.org\/uploads\/images\/201710\/14\/1\/xAuDMxteQy.png?imageView2\/1\/w\/200\/h\/200","introduction":"Ex vel assumenda neque impedit nam.","bound_phone":false,"bound_wechat":false,"last_actived_at":"2019-03-13 15:05:17","created_at":"2019-03-13 15:05:17","updated_at":"2019-04-24 21:06:39"},"category":{"id":2,"name":"\u6559\u7a0b","description":"\u5f00\u53d1\u6280\u5de7\u3001\u63a8\u8350\u6269\u5c55\u5305\u7b49"}}],"meta":{"pagination":{"total":6,"count":6,"per_page":20,"current_page":1,"total_pages":1,"links":null}}}`

        var user models.UserReply
        json.Unmarshal([]byte(aa), &user)                                                                                                                      
                                                                                                                                                               
        u.Data["json"] = user                                                                                                                                  
                                                                                                                                                               
        u.ServeJSON()                                                                                                                                          
}   

// @Title Get
// @Description get user by uid
// @Param       uid             path    string  true            "The key for
// staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid/topics [get]
func (u *UserController) GetUserTopic() {

//        var user models.UserTopic
//        json.Unmarshal([]byte(aa), &user)                                                                                                                      
                                                                                                                                                               
	var topics models.UserTopic
	articles, _ := models.GetAllArticles()

	var topic models.Topic
	for _, article := range articles {
		// beego.Error(article)
		topic.Article = article
		topic.User, _  = models.GetUser(1)
		topic.Category, _  = models.GetCategory(1)
		beego.Error(topic)
		topics.Data = append(topics.Data, topic)
	}

        u.Data["json"] = topics 
                                                                                                                                                               
        u.ServeJSON()                                                                                                                                          

}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router / [put]
func (u *UserController) Put() {
	var scene models.WechatLoginScene
	json.Unmarshal(u.Ctx.Input.RequestBody, &scene)
	beego.Error("=!!!!!!=scene==", scene.User)

	apiurl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wxc4e11081e3d5bdf7&secret=e3284a3123d4ed4e06ee09bf0171bef7&js_code=%s&grant_type=authorization_code", scene.Code)

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	var id models.WechatID

	a.ToJSON(&id)
	beego.Error("openidid === ", id)

	scene.User.BoundWechat = id.Openid
	beego.Error("==!!!!!==", scene)

	models.UpdateUser(scene.User)
	// // models.AddWechatLoginScene(scene)
	u.Data["json"] = id
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
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
func (u *UserController) Login() {
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
