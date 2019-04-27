// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/bbhj/minabbs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/db", beego.NSInclude(&controllers.DBController{})),
		beego.NSNamespace("/api/user", beego.NSInclude( &controllers.UserController{},),),
		beego.NSNamespace("/api/users", beego.NSInclude( &controllers.UserController{},),),
		beego.NSNamespace("/api/categories", beego.NSInclude( &controllers.CategoryController{},),),
		beego.NSNamespace("/api/topics", beego.NSInclude( &controllers.TopicController{},),),
		beego.NSNamespace("/api/weapp/authorizations", beego.NSInclude( &controllers.AuthorizationController{},),),
		beego.NSNamespace("/api/weapp/users", beego.NSInclude( &controllers.UserController{},),),
	)
	beego.AddNamespace(ns)
}
