package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "IndexAbout",
			Router:           `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetComment",
			Router:           `/comment/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetDetails",
			Router:           `/details/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "IndexMessage",
			Router:           `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetReg",
			Router:           `/reg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "IndexReward",
			Router:           `/reward`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["liteblog/controllers:MessageController"],
		beego.ControllerComments{
			Method:           "NewMessage",
			Router:           `/?:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["liteblog/controllers:MessageController"],
		beego.ControllerComments{
			Method:           "Count",
			Router:           `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["liteblog/controllers:MessageController"],
		beego.ControllerComments{
			Method:           "Query",
			Router:           `/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method:           "Del",
			Router:           `/del/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method:           "EditPage",
			Router:           `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:PraiseController"] = append(beego.GlobalControllerRouter["liteblog/controllers:PraiseController"],
		beego.ControllerComments{
			Method:           "Praise",
			Router:           `/:type/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:UesrController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UesrController"],
		beego.ControllerComments{
			Method:           "Activation",
			Router:           `/activation`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("key"),
				param.New("verification"),
			),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:UesrController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UesrController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:UesrController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UesrController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:UesrController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UesrController"],
		beego.ControllerComments{
			Method:           "Reg",
			Router:           `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["liteblog/controllers:UesrController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UesrController"],
		beego.ControllerComments{
			Method:           "SendEmail",
			Router:           `/sendEmail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("email"),
				param.New("key"),
			),
			Params: nil})

}
