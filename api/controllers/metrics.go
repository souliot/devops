package controllers

import (
	"devops/api/models"

	"github.com/gin-gonic/gin"
)

type Metrics struct {
	Base
}

func (c *Metrics) Metrics(ctx *gin.Context) {
	models.Handler.ServeHTTP(ctx.Writer, ctx.Request)
}
