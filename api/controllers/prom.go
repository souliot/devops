package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Prom struct {
	Base
}

// @Tags Prom
// @Summary  生成Prom配置文件
// @Description 生成Prom配置文件
// @Accept  json
// @Produce json
// @Param path                query string false "生成配置文件路径"
// @Param scrapeInterval      query string false "抓取时间间隔"
// @Param evaluationInterval  query string false "报警检测时间间隔"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /prom/conf [get]
func (c *Prom) BuildConfiger(ctx *gin.Context) {
	m := &models.Prom{}
	m.Path = ctx.Query("path")
	m.ScrapeInterval = c.DefaultInt(ctx, "scrapeInterval", 60)
	m.EvaluationInterval = c.DefaultInt(ctx, "evaluationInterval", 60)

	errC, err := m.BuildConfiger()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.RespSuccess)
}

// @Tags Prom
// @Summary  重载Prom配置文件
// @Description 重载Prom配置文件（热更新）
// @Accept  json
// @Produce json
// @Success 200 {object}	resp.Response "返回数据"
// @Router /prom/reload [post]
func (c *Prom) Reload(ctx *gin.Context) {
	m := &models.Prom{}

	errC, err := m.Reload()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.RespSuccess)
}
