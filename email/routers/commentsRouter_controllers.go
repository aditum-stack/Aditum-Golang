package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["email/controllers:EmailInfoController"] = append(beego.GlobalControllerRouter["email/controllers:EmailInfoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/email`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["email/controllers:EmailInfoController"] = append(beego.GlobalControllerRouter["email/controllers:EmailInfoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/email`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["email/controllers:EmailInfoController"] = append(beego.GlobalControllerRouter["email/controllers:EmailInfoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/email/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["email/controllers:EmailInfoController"] = append(beego.GlobalControllerRouter["email/controllers:EmailInfoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/email/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["email/controllers:EmailInfoController"] = append(beego.GlobalControllerRouter["email/controllers:EmailInfoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/email/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
