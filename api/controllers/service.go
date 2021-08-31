package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Base
}

// @Tags 服务
// @Summary  获取服务列表
// @Description 获取服务列表
// @Accept  json
// @Produce json
// @Param env 	query string false "service env"
// @Param path  query string false "service path"
// @Param type 	query string false "service type"
// @Param id    query string false "service id"
// @Success 200 {object}	resp.Response{data=[]ServiceMeta} "返回数据 ServiceMeta"
// @Router /service/all [get]
func (c *Service) All(ctx *gin.Context) {
	req := &models.ServiceRequest{}
	req.Env = ctx.Query("env")
	req.Path = ctx.Query("path")
	req.Typ = ctx.Query("type")
	req.Id = ctx.Query("id")
	req.MetricsType = ctx.Query("metricsType")

	res, errC, err := models.DefaultService.All(req)
	if err != nil {
		ctx.JSON(errC.Status, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(res))
}

// @Tags 服务
// @Summary  删除离线服务节点
// @Description 删除离线服务节点
// @Accept  json
// @Produce json
// @Param env  path string true "service env"
// @Param path path string true "service path"
// @Param type path string true "service type"
// @Param id   path string true "service id"
// @Success 200 {object}	resp.Response
// @Router /service/{env}/{path}/{type}/{id} [delete]
func (c *Service) Delete(ctx *gin.Context) {
	req := &models.ServiceRequest{}
	req.Env = ctx.Param("env")
	req.Path = ctx.Param("path")
	req.Typ = ctx.Param("type")
	req.Id = ctx.Param("id")
	if !c.CheckParams(ctx, req.Env, req.Path, req.Typ, req.Id) {
		return
	}
	errC, err := models.DefaultService.Delete()
	if err != nil {
		ctx.JSON(errC.Status, errC)
		return
	}
	ctx.JSON(200, resp.RespSuccess)
}
