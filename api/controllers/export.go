package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Export struct{}

// @Tags 监控暴露
// @Summary  Export
// @Description 获取用户 Prometheus http_sd_config 的接口
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
