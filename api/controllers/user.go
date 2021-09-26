package controllers

import (
	"devops/api/models"
	"devops/pkg/resp"

	"github.com/gin-gonic/gin"
)

type User struct {
	Base
}

// @Tags 用户
// @Summary  获取用户列表
// @Description 获取用户列表
// @Accept  json
// @Produce json
// @Param name     query string false "node address"
// @Param userName query string false "node address"
// @Param page 		 query string false "page"
// @Param pageSize query string false "pageSize"
// @Success 200 {object}	resp.Response{data=[]models.User} "返回数据 []User"
// @Router /user [get]
func (c *User) All(ctx *gin.Context) {
	m := &models.User{PageQuery: &models.PageQuery{}}
	m.Name = ctx.Query("name")
	m.UserName = ctx.Query("userName")
	m.Page = c.DefaultInt(ctx, "page", 1)
	m.PageSize = c.DefaultInt(ctx, "pageSize", 0)

	exs, errC, err := m.All()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(exs))
}

// @Tags 用户
// @Summary  获取用户信息
// @Description 获取用户信息
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user/{id} [get]
func (c *User) One(ctx *gin.Context) {
	m := new(models.User)
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

// @Tags 用户
// @Summary  添加用户
// @Description 添加用户的接口
// @Accept  json
// @Produce json
// @Param object body models.User true "用户"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user [post]
func (c *User) Add(ctx *gin.Context) {
	m := new(models.User)
	m.Roles = make([]string, 0)
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

// @Tags 用户
// @Summary  修改用户
// @Description 修改用户的接口
// @Accept  json
// @Produce json
// @Param object body models.User true "用户"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user [put]
func (c *User) Update(ctx *gin.Context) {
	m := new(models.User)
	m.Roles = make([]string, 0)
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

// @Tags 用户
// @Summary  修改用户
// @Description 修改用户的接口
// @Accept  json
// @Produce json
// @Param object body models.User true "用户"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user/password [put]
func (c *User) UpdatePassword(ctx *gin.Context) {
	m := new(models.User)
	m.Roles = make([]string, 0)
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}

	errC, err := m.UpdatePassword()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}

// @Tags 用户
// @Summary  删除用户
// @Description 删除用户的接口
// @Accept  json
// @Produce json
// @Param id path string true "node id"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user/{id} [delete]
func (c *User) Delete(ctx *gin.Context) {
	m := new(models.User)
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

// @Tags 用户
// @Summary  用户登录
// @Description 用户登录的接口
// @Accept  json
// @Produce json
// @Param object body models.User true "用户"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user/login [post]
func (c *User) Login(ctx *gin.Context) {
	m := new(models.User)
	m.Roles = make([]string, 0)
	if err := ctx.ShouldBind(m); err != nil {
		c.CheckError(err, ctx)
		return
	}

	token, errC, err := m.Login()
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(token))
}

// @Tags 用户
// @Summary  用户登出
// @Description 用户登出的接口
// @Accept  json
// @Produce json
// @Param object body models.User true "用户"
// @Success 200 {object}	resp.Response "返回数据 User"
// @Router /user/logout [post]
func (c *User) Logout(ctx *gin.Context) {
	m := new(models.User)
	token := ctx.GetHeader("Authorization")
	errC, err := m.Logout(token)
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.RespSuccess)
}

// @Tags 用户
// @Summary  获取登录用户信息
// @Description 获取登录用户信息的接口
// @Accept  json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object}	resp.Response{data=models.User} "返回数据 User"
// @Router /user/getUserInfo [get]
func (c *User) GetUserInfo(ctx *gin.Context) {
	m := new(models.User)
	token := ctx.GetHeader("Authorization")
	errC, err := m.GetUserInfo(token)
	if err != nil {
		ctx.JSON(200, errC)
		return
	}
	ctx.JSON(200, resp.NewSuccess(m))
}
