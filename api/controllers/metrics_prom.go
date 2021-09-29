package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type MetricsProm struct {
	Base
}

// @Tags MetricsProm
// @Summary  获取主机监控信息
// @Description 获取主机监控
// @Accept  json
// @Produce json
// @Success 200 {object}	resp.Response "返回数据"
// @Router /metricsProm/hostinfo [get]
func (c *MetricsProm) HostInfo(ctx *gin.Context) {
	m := &models.MetricsProm{}
	res, errC, err := m.HostInfo()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(res))
}
