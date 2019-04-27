package models

import (
	_ "errors"
	_ "github.com/jinzhu/gorm"
	_ "github.com/astaxie/beego"
)

var ()

func AddArticle(a Article) (err error) {
	conn.FirstOrCreate(&a, Article{Title: a.Title})
	// conn.Save(&a)
	return
}

func GetAllArticles() (c []Article, err error) {
	pagesize := 5
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

func  TestRelated() {
	var a  Article
	conn.First(&a)
	// conn.Debug().Model(&a).Related(&a.ID).Find(&a.User)
}
