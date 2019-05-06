package models

import (
	_ "github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "strconv"
	_ "strings"
	"time"
)

type (
	Babyinfo struct {
		gorm.Model
		// ID        string `gorm:"primary_key"`
		//  Timestamp int64
		// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
		UUID      string
		AvatarUrl string
		Nickname  string
		// 0: unkown,  1: male   2: female
		Gender uint
		Title           string
		Subject         string
		Characters      string
		Details         string
		DataFrom        string
		BirthedProvince string
		BirthedCity     string
		BirthedCountry  string
		BirthedAddress  string
		BirthedAt       time.Time `gorm:"type:datetime"`

		MissedCountry  string
		MissedProvince string
		MissedCity     string
		MissedAddress  string
		MissedAt       time.Time `gorm:"column:missed_at;type:datetime"`
		Handler        string
		Babyid         string
		Category       string
		Height         string
		SyncStatus     int	`gorm:"column:syncstatus;default:0"`
	}

)

func GetAllBabyinfo() (data []Babyinfo, err error) {
	pagesize := 10
	conn.Limit(pagesize).Find(&data)
	return
}
