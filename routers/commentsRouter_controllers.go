package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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
            Method: "Message",
            Router: "/message",
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

    beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Index",
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
            Method: "GoRegister",
            Router: "/goRegister",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
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
            Method: "GoPlay",
            Router: "/play",
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

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: "/upload",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserVideos",
            Router: "/userVideos",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
