package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Environment struct {
	Base
}

// @Tags 环境
// @Summary  获取环境列表
// @Description 获取环境列表
// @Accept  json
// @Produce json
// @Param name 	   query string false "node type"
// @Param page 		 query string false "page"
// @Param pageSize query string false "pageSize"
// @Success 200 {object}	resp.Response{data=[]models.Environment} "返回数据 []Environment"
// @Router /env [get]
func (c *Environment) All(ctx *gin.Context) {
	m := &models.Environment{PageQuery: &models.PageQuery{}}
	m.Name = ctx.Query("name")
	m.Page = c.DefaultInt(ctx, "page", 1)
	m.PageSize = c.DefaultInt(ctx, "pageSize", 0)

	exs, errC, err := m.All()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(exs))
}

// @Tags 环境
// @Summary  获取环境信息
// @Description 获取环境信息
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.Environment} "返回数据 Environment"
// @Router /env/{id} [get]
func (c *Environment) One(ctx *gin.Context) {
	m := new(models.Environment)
	m.Id = ctx.Param("id")
	if !c.CheckParams(ctx, m.Id) {
		return
	}
	errC, err := m.One()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}

// @Tags 环境
// @Summary  添加环境
// @Description 添加环境的接口
// @Accept  json
// @Produce json
// @Param object body models.Environment true "环境节点"
// @Success 200 {object}	resp.Response{data=models.Environment} "返回数据 Environment"
// @Router /env [post]
func (c *Environment) Add(ctx *gin.Context) {
	m := new(models.Environment)
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

// @Tags 环境
// @Summary  修改环境
// @Description 修改环境的接口
// @Accept  json
// @Produce json
// @Param object body models.Environment true "环境节点"
// @Success 200 {object}	resp.Response{data=models.Environment} "返回数据 Environment"
// @Router /env [put]
func (c *Environment) Update(ctx *gin.Context) {
	m := new(models.Environment)
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}

	errC, err := m.Update()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}

// @Tags 环境
// @Summary  删除环境
// @Description 删除环境的接口
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.Environment} "返回数据 Environment"
// @Router /env/{id} [delete]
func (c *Environment) Delete(ctx *gin.Context) {
	m := new(models.Environment)
	m.Id = ctx.Param("id")
	if !c.CheckParams(ctx, m.Id) {
		return
	}
	errC, err := m.Delete()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
