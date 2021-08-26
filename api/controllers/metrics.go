package controllers

import (
	"devops/api/models"

	"github.com/gin-gonic/gin"
)

type Metrics struct{}

// @Tags 监控暴露
// @Summary  Export
// @Description 获取用户 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce text
// @Success 200 string	resp.Response{data=models.Export}
// @Router /metrics [get]
func (c *Metrics) Metrics(ctx *gin.Context) {
	models.Handler.ServeHTTP(ctx.Writer, ctx.Request)
}
