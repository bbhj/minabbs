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
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:AuthorizationController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:AuthorizationController"],
        beego.ControllerComments{
            Method: "Token",
            Router: `/token`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"],
        beego.ControllerComments{
            Method: "InitData",
            Router: `/data`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"],
        beego.ControllerComments{
            Method: "InitDB",
            Router: `/db`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"],
        beego.ControllerComments{
            Method: "Droptable",
            Router: `/droptable`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"],
        beego.ControllerComments{
            Method: "Run",
            Router: `/run`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:InitController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:LostController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:LostController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:LostController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:LostController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/detail/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:LostController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:LostController"],
        beego.ControllerComments{
            Method: "Search",
            Router: `/search`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetReply",
            Router: `/:id/replies`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "PostReply",
            Router: `/:uid/replies`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserReply",
            Router: `/:uid/replies`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserTopic",
            Router: `/:uid/topics`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:VolunteerController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:VolunteerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:VolunteerController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:VolunteerController"],
        beego.ControllerComments{
            Method: "CheckVolunteer",
            Router: `/check`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:WechatController"] = append(beego.GlobalControllerRouter["github.com/bbhj/minabbs/controllers:WechatController"],
        beego.ControllerComments{
            Method: "CreateQRcode",
            Router: `/createqrcode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
