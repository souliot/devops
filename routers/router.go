package routers

import (
	"common/models"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	models.DefaultMetrics.Init()
	beego.Handler("/", models.Handler)
	beego.Handler("/metrics", models.Handler)
}
