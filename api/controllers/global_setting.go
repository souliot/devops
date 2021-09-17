package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type GlobalSetting struct {
	Base
}

// @Tags GlobalSetting
// @Summary  获取全局配置
// @Description 获取全局配置
// @Accept  json
// @Produce json
// @Param env  path string false "环境"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /globalSetting/:env [get]
func (c *GlobalSetting) GetGlobalSetting(ctx *gin.Context) {
	m := &models.GlobalSetting{}
	m.Env = ctx.Param("env")
	if !c.CheckParams(ctx, m.Env) {
		return
	}

	res, errC, err := m.GetGlobalSetting()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(res))
}

// @Tags GlobalSetting
// @Summary  修改全局配置
// @Description 修改全局配置
// @Accept  json
// @Produce json
// @Param env    query string true "环境名称"
// @Param object body models.GlobalSetting true "全局配置"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /globalSetting [put]
func (c *GlobalSetting) SetGlobalSetting(ctx *gin.Context) {
	m := &models.GlobalSetting{}
	m.Env = ctx.Query("env")
	if !c.CheckParams(ctx, m.Env) {
		return
	}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		errC := resp.ErrUserInput
		ctx.JSON(200, errC)
		return
	}
	errC, err := m.SetGlobalSetting(data)
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
