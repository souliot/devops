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
// @Summary  获取监控节点列表
// @Description 获取 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Param type 	   query string false "node type"
// @Param address  query string false "node address"
// @Param page 		 query string false "page"
// @Param pageSize query string false "pageSize"
// @Success 200 {object}	resp.Response{data=[]models.Export}
// @Router /export [get]
func (c *Export) All(ctx *gin.Context) {
	m := &models.Export{PageQuery: &models.PageQuery{}}
	m.Type = ctx.Query("type")
	m.Address = ctx.Query("address")
	m.Page = c.DefaultInt(ctx, "page", 1)
	m.PageSize = c.DefaultInt(ctx, "pageSize", 0)

	exs, errC, err := m.All()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(exs))
}

// @Tags 监控
// @Summary  获取监控节点信息
// @Description 获取 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.Export}
// @Router /export/{id} [get]
func (c *Export) One(ctx *gin.Context) {
	m := new(models.Export)
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

// @Tags 监控
// @Summary  获取监控节点
// @Description 获取 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Param type path string true "node type"
// @Success 200 {object}	resp.Response{data=models.Export}
// @Router /export/type/{type} [get]
func (c *Export) Node(ctx *gin.Context) {
	m := new(models.Export)
	m.Type = ctx.Param("type")
	if !c.CheckParams(ctx, m.Type) {
		return
	}
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

// @Tags 监控
// @Summary  修改监控节点
// @Description 修改 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Param object body models.Export true "监控节点"
// @Success 200 {object}	resp.Response{data=models.Export}
// @Router /export [put]
func (c *Export) Update(ctx *gin.Context) {
	m := new(models.Export)
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

// @Tags 监控
// @Summary  删除监控节点
// @Description 删除 Prometheus http_sd_config 的接口
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.Export}
// @Router /export/{id} [delete]
func (c *Export) Delete(ctx *gin.Context) {
	m := new(models.Export)
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
