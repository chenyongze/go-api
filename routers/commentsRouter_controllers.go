package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:AuctionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:CmsArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:DingController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:DingController"],
		beego.ControllerComments{
			Method: "Reids_hash_xinxi",
			Router: `/redis_h`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:DingController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:DingController"],
		beego.ControllerComments{
			Method: "Redis",
			Router: `/sh-redis`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:LotsController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:LotsController"],
		beego.ControllerComments{
			Method: "Getlots",
			Router: `/getlots`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:LotsController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:LotsController"],
		beego.ControllerComments{
			Method: "LotInfo",
			Router: `/lotinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:LotsController"] = append(beego.GlobalControllerRouter["github.com/chenyongze/go-api/controllers:LotsController"],
		beego.ControllerComments{
			Method: "Todo",
			Router: `/todo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
