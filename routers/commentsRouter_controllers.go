package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["common/controllers:ExportController"] = append(beego.GlobalControllerRouter["common/controllers:ExportController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/export/node",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
