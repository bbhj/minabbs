package main

import (
	"github.com/bbhj/minabbs/models"
	_ "github.com/bbhj/minabbs/routers"

	"github.com/astaxie/beego"
	"os"
	"path"
	"time"
)

func main() {
	beego.SetLogFuncCall(true)
	logpath := path.Join(beego.AppPath, "logs")
	os.Mkdir(logpath, 0755)
	logfile := path.Join(logpath, beego.BConfig.AppName+".log")
	beego.SetLogger("file", `{"filename":"`+logfile+`", "level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	models.StartTime = time.Now()
	models.Connect()
	defer models.Close()
	beego.Info(beego.BConfig.AppName, "start...")
	beego.Info("runmode:", beego.BConfig.RunMode)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
