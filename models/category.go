package models

import (
	_ "errors"
	_ "github.com/jinzhu/gorm"
)

var ()

func AddCategory(c Category) (err error) {
	conn.FirstOrCreate(&c, Category{Name: c.Name})
	return
}
func GetAllCategories() (c []Category, err error) {
	pagesize := 5
	conn.Limit(pagesize).Find(&c)
	return
}


func GetCategory(id int) (c Category, err error) {
	conn.First(&c, id)
	if ( conn.Error != nil && conn.RowsAffected > 0 ) {
	 	err = nil
	}
	return
}
