package models

import (
	_ "errors"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"time"
)

func init() {
}

type (
	Profile struct {
		Openid      string
		AccessToken string
		UserID      uint
	}
	// Openid  string `json:"Openid"`
	// AccessToken string `json:"AccessToken"`
	// UserID	uint	`json:"UserID"`
	// AccessToken string `json:"access_token"`
	// UserID	uint	`json:"user_id"`

	UserReply struct {
		Data []Article `json:"data"`
		Meta `json:"meta"`
	}

	UserTopic struct {
		Data []Topic `json:"data"`
		Meta `json:"meta"`
	}

	Topic struct {
		Article
		Category   `json:"category"`
		User       User    `json:"user"`
		TopReplies Replies `json:"topReplies"`
	}

	// TopReplies      Replies `json:"topReplies"`

	Meta struct {
		Pagination struct {
			Count       int         `json:"count"`
			CurrentPage int         `json:"current_page"`
			Links       interface{} `json:"links"`
			PerPage     int         `json:"per_page"`
			Total       int         `json:"total"`
			TotalPages  int         `json:"total_pages"`
		} `json:"pagination"`
	}

	// LastActivedAt string `json:"last_actived_at"`
	User struct {
		gorm.Model
		Avatar        string    `json:"avatar"`
		BoundPhone    string    `json:"phone"`
		BoundWechat   string    `json:"openid"`
		Email         string    `json:"email"`
		Introduction  string    `json:"introduction"`
		LastActivedAt time.Time `sql:"DEFAULT:current_timestamp"`
		Name          string    `json:"name"`
		NickName      string    `json:"nickName"`
	}
	// CreatedAt     string `json:"created_at"`
	// ID            int    `json:"id"`
	// UpdatedAt     string `json:"updated_at"`
)

func AddUser(u User) (err error) {
	// conn.FirstOrCreate(&u, User{Email: u.Email})
	conn.FirstOrCreate(&u, User{Avatar: u.Avatar})
	// conn.Save(&u)
	return
}

func UpdateUser(u User) (user User, err error) {
	conn.Debug().Model(&user).Select("name", "nick_name", "avatar").Where("bound_wechat = ?", u.BoundWechat).Update(&u)
	beego.Info(conn.Error)
	return
}

func AddUserByWechat(openid string) (user User, err error) {
	conn.FirstOrCreate(&user, User{BoundWechat: openid})
	return
}

func GetAllUsers() (u []User, err error) {
	pagesize := 5
	conn.Limit(pagesize).Find(&u)
	return
}

func GetUser(uid int) (u User, err error) {
	conn.First(&u, uid)
	if conn.Error != nil && conn.RowsAffected > 0 {
		err = nil
	}
	return
}
