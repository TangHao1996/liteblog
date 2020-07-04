package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["liteblog/controllers:CommentController"] = append(beego.GlobalControllerRouter["liteblog/controllers:CommentController"],
        beego.ControllerComments{
            Method: "New",
            Router: "/new/:notekey",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "About",
            Router: "/about",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/authentication",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Details",
            Router: "/details/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GoRegister",
            Router: "/goRegister",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GoLogin",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Message",
            Router: "/message",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GoPlay",
            Router: "/play",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "User",
            Router: "/user",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:LikeController"] = append(beego.GlobalControllerRouter["liteblog/controllers:LikeController"],
        beego.ControllerComments{
            Method: "Like",
            Router: "/:type/:key",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "NewNote",
            Router: "/new",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/save/:key",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "UploadPage",
            Router: "/upload",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
