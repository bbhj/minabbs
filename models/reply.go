package models

import (
	_ "errors"
	_ "github.com/jinzhu/gorm"
	_ "github.com/astaxie/beego"
)

var ()

func AddReply(a Reply) (err error) {
	conn.Save(&a)
	return
}

func GetAllReplies() (c []Reply, err error) {
	pagesize := 5
	conn.Limit(pagesize).Find(&c)
	return
}


func GetReply(id int) (c Reply, err error) {
	conn.First(&c, id)
	if ( conn.Error != nil && conn.RowsAffected > 0 ) {
	 	err = nil
	}
	return
}
