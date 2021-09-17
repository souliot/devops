package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type TypSetting struct {
	Base
}

// @Tags TypSetting
// @Summary  获取服务类型配置
// @Description 获取服务类型配置
// @Accept  json
// @Produce json
// @Param env  path string false "环境"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /typSetting/:env/:typ [get]
func (c *TypSetting) GetTypSetting(ctx *gin.Context) {
	m := &models.TypSetting{}
	m.Env = ctx.Param("env")
	m.Typ = ctx.Param("typ")
	if !c.CheckParams(ctx, m.Env, m.Typ) {
		return
	}

	res, errC, err := m.GetTypSetting()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(res))
}

// @Tags TypSetting
// @Summary  修改服务类型配置
// @Description 修改服务类型配置
// @Accept  json
// @Produce json
// @Param env    query string true "环境名称"
// @Param object body models.TypSetting true "服务类型配置"
// @Success 200 {object}	resp.Response "返回数据"
// @Router /typSetting [put]
func (c *TypSetting) SetTypSetting(ctx *gin.Context) {
	m := &models.TypSetting{}
	m.Env = ctx.Query("env")
	m.Typ = ctx.Query("typ")
	if !c.CheckParams(ctx, m.Env, m.Typ) {
		return
	}

	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		errC := resp.ErrUserInput
		ctx.JSON(200, errC)
		return
	}
	errC, err := m.SetTypSetting(data)
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
