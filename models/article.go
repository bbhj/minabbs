package models

import (
	_ "errors"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type (
	Article1 struct {
		gorm.Model
		Body            string `json:"body"`
		Category        `json:"category"`
		CategoryID      int         `json:"category_id"`
		Excerpt         string      `json:"excerpt"`
		LastReplyUserID int         `json:"last_reply_user_id"`
		ReplyCount      int         `json:"reply_count"`
		Slug		string 	    `json:"slug"`
		Title           string      `json:"title"`
		User		User	`gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`
		UserID          int `json:"user_id"`
		ViewCount       int `json:"view_count"`
		TopReplies      Replies `json:"topReplies"`
	}

	Article struct {
		gorm.Model
		Body            string `json:"body"`
		CategoryID      int         `json:"category_id"`
		Excerpt         string      `json:"excerpt"`
		LastReplyUserID int         `json:"last_reply_user_id"`
		ReplyCount      int         `json:"reply_count"`
		Slug		string 	    `json:"slug"`
		Title           string      `json:"title"`
		UserID          int `json:"user_id"`
		ViewCount       int `json:"view_count"`
	}

		// TopReplies      Replies `json:"topReplies"`
	Article2 struct {
		Article
		Category        `json:"category"`
		User		User	`json:"user"`
		TopReplies      Replies `json:"topReplies"`
	}
)

func AddArticle(a Article) (article Article, err error) {
	// conn.Debug().FirstOrCreate(&a, Article{Title: a.Title})
	conn.Save(&a)
	return
}

func GetAllArticles() (c []Article, err error) {
	pagesize := 10000
	conn.Limit(pagesize).Find(&c)
	return
}


func GetArticle(id int) (c Article, err error) {
	beego.Error("this id =====", id)
	conn.First(&c, id)
	UpdateArticleView(id)
	if ( conn.Error != nil && conn.RowsAffected > 0 ) {
	 	err = nil
	}
	return
}

// Update article reply count
func UpdateArticleReply(topic_id, user_id int) (a Article, err error) {
	// conn.Table("article_summaries").Where("babyid = ?", babyid).UpdateColumn("Forward", gorm.Expr("forward + ?", 1))
	conn.Debug().Model(&a).Where("id = ?", topic_id).UpdateColumn("last_reply_user_id", user_id).UpdateColumn("reply_count", gorm.Expr("reply_count + ?", 1))
	return
}

// update article view count
func UpdateArticleView(topic_id int) (a Article, err error) {
	conn.Debug().Model(&a).Where("id = ?", topic_id).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1))
	return
}

func  TestRelated() {
	var a  Article
	conn.First(&a)
	// conn.Debug().Model(&a).Related(&a.ID).Find(&a.User)
}
