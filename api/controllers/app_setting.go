package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type AppSetting struct {
	Base
}

// @Tags AppSetting
// @Summary  获取服务实例配置
// @Description 获取服务实例配置
// @Accept  json
// @Produce json
// @Param env  path string false "环境"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /appSetting/:env/:typ/:id [get]
func (c *AppSetting) GetAppSetting(ctx *gin.Context) {
	m := &models.AppSetting{}
	m.Env = ctx.Param("env")
	m.Typ = ctx.Param("typ")
	m.Id = ctx.Param("id")
	if !c.CheckParams(ctx, m.Env, m.Typ) {
		return
	}

	res, errC, err := m.GetAppSetting()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(res))
}

// @Tags AppSetting
// @Summary  修改服务实例配置
// @Description 修改服务实例配置
// @Accept  json
// @Produce json
// @Param env    query string true "环境名称"
// @Param object body models.AppSetting true "服务实例配置"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /appSetting [put]
func (c *AppSetting) SetAppSetting(ctx *gin.Context) {
	m := &models.AppSetting{}
	m.Env = ctx.Query("env")
	m.Typ = ctx.Query("typ")
	m.Id = ctx.Query("id")
	if !c.CheckParams(ctx, m.Env, m.Typ) {
		return
	}

	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		errC := resp.ErrUserInput
		ctx.JSON(200, errC)
		return
	}
	errC, err := m.SetAppSetting(data)
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
