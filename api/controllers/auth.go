package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
	logs "github.com/souliot/siot-log"
)

type Auth struct{}

func (c *Auth) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logs.Info("接口调用：", ctx.Request.RequestURI)
		valid := false
		token := ctx.GetHeader("Authorization")
		if token != "" {
			t := &models.Token{
				Token: token,
			}

			valid = t.Valid()
			if valid {
				return
			}
			t.Delete()
		}

		appid := ctx.GetHeader("Appid")
		secret := ctx.GetHeader("Secret")
		if appid == "" || secret == "" {
			ctx.Abort()
			errC := resp.ErrTokenInValid
			ctx.JSON(200, errC)
			return
		}
		u := &models.AppUser{
			Appid:  appid,
			Secret: secret,
		}
		valid = u.Valid()
		if !valid {
			ctx.Abort()
			errC := resp.ErrAppid
			ctx.JSON(200, errC)
			return
		}
		return
	}
}
