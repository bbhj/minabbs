package models

import (
	_ "errors"
	"github.com/jinzhu/gorm"
)

type (
	Categories struct {
		Data []Category `json:"data"`
	}

	Category struct {
		gorm.Model
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

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
	if conn.Error != nil && conn.RowsAffected > 0 {
		err = nil
	}
	return
}
