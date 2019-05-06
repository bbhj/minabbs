package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

func Connect() (db *gorm.DB, err error) {
	conn, err = gorm.Open("mysql", beego.AppConfig.String("gdbc"))
	if ( err != nil ) {
		beego.Error("Can't connect database, db info: ", beego.AppConfig.String("gdbc"))
	}
	
	return
}

func Close() {
	conn.Close()
}

func Init() (flag bool, errmsg string) {
	
	conn = conn.Set("gorm:table_options", "CHARSET=utf8mb4")
	
	conn = conn.AutoMigrate(&User{})
	conn = conn.AutoMigrate(&Category{})
	conn = conn.AutoMigrate(&Article{})
	conn = conn.AutoMigrate(&Reply{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Volunteer{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&WechatLoginScene{})

	beego.Error("init db in pool", conn)
	if (conn.Error == nil){
		flag = true
	} else {
		errmsg = fmt.Sprintf("%s", conn.Error)
	}
	beego.Info("models init info, status:", flag, ", errmsg:", conn.Error )
	return
}

func Droptable() (flag bool, errmsg string) {
	conn.DropTable(&Article{})
	conn.DropTable(&Reply{})
	return
}
