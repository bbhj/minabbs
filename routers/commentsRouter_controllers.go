package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:AuthorizationController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:AuthorizationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:DBController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:DBController"],
		beego.ControllerComments{
			Method: "InitDB",
			Router: `/init`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:DBController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:DBController"],
		beego.ControllerComments{
			Method: "InitData",
			Router: `/initdata`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:DBController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:DBController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "PostReply",
			Router: `/:uid/replies`,
			AllowHTTPMethods: []string{"Post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "GetReply",
			Router: `/:uid/replies`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserReply",
			Router: `/:uid/replies`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserTopic",
			Router: `/:uid/topics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
