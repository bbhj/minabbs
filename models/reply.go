package models

import (
	_ "errors"
	"github.com/jinzhu/gorm"
	_ "github.com/astaxie/beego"
)

type (
	Reply struct {
		gorm.Model
		Content   string `json:"content"`
		TopicID   int    `json:"topic_id"`
		UserID    int `json:"user_id"`
	}
)

func AddReply(a Reply) (err error) {
	conn.Save(&a)
	UpdateArticleReply(a.TopicID, a.UserID)
	return
}

func GetAllReplies() (c []Reply, err error) {
	pagesize := 5
	conn.Limit(pagesize).Find(&c)
	return
}

func GetRepliesByTopicID(topic_id int) (replies []Reply, err error) {
	pagesize := 10
	conn.Where("topic_id = ?", topic_id).Limit(pagesize).Find(&replies)
	return
}

func GetReply(id int) (c Reply, err error) {
	conn.First(&c, id)
	if ( conn.Error != nil && conn.RowsAffected > 0 ) {
	 	err = nil
	}
	return
}
