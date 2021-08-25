package controllers

import (
	"common/models"

	beego "github.com/beego/beego/v2/server/web"
)

type ExportController struct {
	beego.Controller
}

// @router /export/node [get]
func (c *ExportController) Get() {
	ex := &models.Export{}
	exs := ex.Node()
	c.Data["json"] = exs
	c.ServeJSON()
	return
}
