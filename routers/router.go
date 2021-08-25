package routers

import (
	"common/controllers"
	"common/models"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	models.DefaultMetrics.Init()
	beego.Handler("/", models.Handler)
	beego.Handler("/metrics", models.Handler)
	beego.Router("/export/node", &controllers.ExportController{})
}
