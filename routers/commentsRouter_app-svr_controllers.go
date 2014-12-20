package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["app-svr/controllers:RegisterController"] = append(beego.GlobalControllerRouter["app-svr/controllers:RegisterController"],
		beego.ControllerComments{
			"SmsCode",
			`/register/smscode`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["app-svr/controllers:RegisterController"] = append(beego.GlobalControllerRouter["app-svr/controllers:RegisterController"],
		beego.ControllerComments{
			"SmsCodeAtk",
			`/register/smscode/atk`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["app-svr/controllers:RegisterController"] = append(beego.GlobalControllerRouter["app-svr/controllers:RegisterController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["app-svr/controllers:LoginController"] = append(beego.GlobalControllerRouter["app-svr/controllers:LoginController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["app-svr/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["app-svr/controllers:MemberCardController"],
		beego.ControllerComments{
			"BindCard",
			`/card/:card/bind/:owner`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["app-svr/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["app-svr/controllers:MemberCardController"],
		beego.ControllerComments{
			"UnBindCard",
			`/card/:card/unbind`,
			[]string{"post"},
			nil})

}
