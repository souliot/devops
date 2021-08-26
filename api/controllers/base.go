package controllers

import (
	"devops/pkg/resp"
	"devops/pkg/trans"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Base struct{}

func (c *Base) CheckError(err error, ctx *gin.Context) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		errC := resp.ErrUserInput
		errC.MoreInfo = err.Error()
		ctx.JSON(200, errC)
		return
	}
	errC := resp.ErrUserInput
	errC.MoreInfo = removeTopStruct(errs.Translate(trans.Trans))
	ctx.JSON(200, errC)
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
