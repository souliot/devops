package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type PromJob struct {
	Base
}

// @Tags 任务
// @Summary  获取任务列表
// @Description 获取任务列表
// @Accept  json
// @Produce json
// @Param jobName  query string false "node address"
// @Param page 		 query string false "page"
// @Param pageSize query string false "pageSize"
// @Success 200 {object}	resp.Response{data=[]models.PromJob} "返回数据 []PromJob"
// @Router /promjob [get]
func (c *PromJob) All(ctx *gin.Context) {
	m := &models.PromJob{PageQuery: &models.PageQuery{}}
	m.JobName = ctx.Query("jobName")
	m.Page = c.DefaultInt(ctx, "page", 1)
	m.PageSize = c.DefaultInt(ctx, "pageSize", 0)

	exs, errC, err := m.All()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(exs))
}

// @Tags 任务
// @Summary  获取任务信息
// @Description 获取任务信息
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.PromJob} "返回数据 PromJob"
// @Router /promjob/{id} [get]
func (c *PromJob) One(ctx *gin.Context) {
	m := new(models.PromJob)
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

// @Tags 任务
// @Summary  添加任务
// @Description 添加任务的接口
// @Accept  json
// @Produce json
// @Param object body models.PromJob true "任务节点"
// @Success 200 {object}	resp.Response{data=models.PromJob} "返回数据 PromJob"
// @Router /promjob [post]
func (c *PromJob) Add(ctx *gin.Context) {
	m := new(models.PromJob)
	m.Targets = make([]string, 0)
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}

	if len(m.Targets) <= 0 && m.Url == "" {
		errC := resp.ErrUserInput
		errC.MoreInfo = "targets 跟 url 不能同时为空！"
		ctx.JSON(200, errC)
		return
	}

	errC, err := m.Add()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}

// @Tags 任务
// @Summary  修改任务
// @Description 修改任务的接口
// @Accept  json
// @Produce json
// @Param object body models.PromJob true "任务节点"
// @Success 200 {object}	resp.Response{data=models.PromJob} "返回数据 PromJob"
// @Router /promjob [put]
func (c *PromJob) Update(ctx *gin.Context) {
	m := new(models.PromJob)
	m.Targets = make([]string, 0)
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}

	if len(m.Targets) <= 0 && m.Url == "" {
		errC := resp.ErrUserInput
		errC.MoreInfo = "targets 跟 url 不能同时为空！"
		ctx.JSON(200, errC)
		return
	}

	errC, err := m.Update()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}

// @Tags 任务
// @Summary  删除任务
// @Description 删除任务的接口
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.PromJob} "返回数据 PromJob"
// @Router /promjob/{id} [delete]
func (c *PromJob) Delete(ctx *gin.Context) {
	m := new(models.PromJob)
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
