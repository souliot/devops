package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Base
}

// @Tags Controller
// @Summary  修改服务实例配置
// @Description 修改服务实例配置
// @Accept  json
// @Produce json
// @Param env    query string true "环境名称"
// @Param object body models.Controller true "服务实例配置"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /controller [put]
func (c *Controller) Controller(ctx *gin.Context) {
	m := &models.Controller{}
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}
	if !c.CheckParams(ctx, m.Env, m.Id) {
		return
	}

	errC, err := m.Controller()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
