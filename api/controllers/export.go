package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Export struct {
	Base
}

// @Tags 监控
// @Summary  获取监控节点
// @Description 获取 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Success 200 {object}	resp.Response{data=models.Export}
// @Router /export/node [get]
func (c *Export) Node(ctx *gin.Context) {
	m := new(models.Export)
	ex, errC, err := m.Node()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(ex))
}

// @Tags 监控
// @Summary  添加监控节点
// @Description 添加 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Param object body models.Export true "监控节点"
// @Success 200 {object}	resp.Response{data=models.Export}
// @Router /export [post]
func (c *Export) Add(ctx *gin.Context) {
	m := new(models.Export)
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}

	errC, err := m.Add()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
