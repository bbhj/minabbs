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
	"github.com/astaxie/beego/context"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		// _, ok := ctx.Input.Session("uid").(int)
		// if !ok && ctx.Request.RequestURI != "/login" {
		// 	ctx.Redirect(302, "/login")
		// }

		// et := jwtbeego.EasyToken{}
		// valido, _, _ := et.ValidateToken(tokenString)
		// if !valido {
		// 	c.Ctx.Output.SetStatus(401)
		// 	c.Data["json"] = "Permission Deny"
		// 	c.ServeJSON()
		// }
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/init", beego.NSInclude(&controllers.InitController{})),
		beego.NSNamespace("/api/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/api/users", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/api/category", beego.NSInclude(&controllers.CategoryController{})),
		beego.NSNamespace("/api/lost", beego.NSInclude(&controllers.LostController{})),
		beego.NSNamespace("/api/topics", beego.NSInclude(&controllers.TopicController{})),
		beego.NSNamespace("/api/weapp/authorizations", beego.NSInclude(&controllers.AuthorizationController{})),
		beego.NSNamespace("/api/weapp/users", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/api/wechat", beego.NSInclude(&controllers.WechatController{})),
	)
	beego.AddNamespace(ns)
}
