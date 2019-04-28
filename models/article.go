package models

import (
	_ "errors"
	"github.com/jinzhu/gorm"
	_ "github.com/astaxie/beego"
)

var ()

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
	conn.First(&c, id)
	if ( conn.Error != nil && conn.RowsAffected > 0 ) {
	 	err = nil
	}
	return
}

func UpdateArticleReply(topic_id, user_id int) (a Article, err error) {
	// conn.Table("article_summaries").Where("babyid = ?", babyid).UpdateColumn("Forward", gorm.Expr("forward + ?", 1))
	conn.Debug().Model(&a).Where("id = ?", topic_id).UpdateColumn("last_reply_user_id", user_id).UpdateColumn("reply_count", gorm.Expr("reply_count + ?", 1))
	return
}


func  TestRelated() {
	var a  Article
	conn.First(&a)
	// conn.Debug().Model(&a).Related(&a.ID).Find(&a.User)
}
