package controllers

import (
	"devops/pkg/resp"
	"devops/pkg/trans"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Base struct{}

func (c *Base) CheckParams(ctx *gin.Context, params ...string) bool {
	for _, v := range params {
		if v == "" {
			errC := resp.ErrUserInput
			errC.MoreInfo = "参数不能为空"
			ctx.JSON(200, errC)
			return false
		}
	}
	return true
}

func (c *Base) DefaultBool(ctx *gin.Context, key string, defaultValue bool) (b bool) {
	v := ctx.Query(key)
	if v == "" {
		return defaultValue
	}
	b, _ = strconv.ParseBool(v)
	return
}

func (c *Base) DefaultInt(ctx *gin.Context, key string, defaultValue int) (i int) {
	v := ctx.Query(key)
	if v == "" {
		return defaultValue
	}
	i, _ = strconv.Atoi(v)
	return
}

func (c *Base) DefaultInt32(ctx *gin.Context, key string, defaultValue int32) (i int32) {
	v := ctx.Query(key)
	if v == "" {
		return defaultValue
	}
	d, _ := strconv.Atoi(v)
	return int32(d)
}

func (c *Base) DefaultInt64(ctx *gin.Context, key string, defaultValue int64) (i int64) {
	v := ctx.Query(key)
	if v == "" {
		return defaultValue
	}
	i, _ = strconv.ParseInt(v, 10, 64)
	return
}

func (c *Base) HandlerNoRouter(ctx *gin.Context) {
	errC := resp.Err404
	errC.MoreInfo = "访问页面不存！"
	ctx.JSON(200, errC)
	return
}

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
