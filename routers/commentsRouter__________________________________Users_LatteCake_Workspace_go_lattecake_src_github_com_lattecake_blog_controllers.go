package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:HomeController"],
		beego.ControllerComments{
			Method: "About",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Reward",
			Router: `/reward`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/learn`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/lattecake/blog/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/post/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
