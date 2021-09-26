package models

import (
	"devops/pkg/auth"
	"devops/pkg/resp"
	"errors"
	"fmt"
	"time"
)

var (
	ErrUserExist = fmt.Errorf("用户名已存在！")
)

type User struct {
	*PageQuery `orm:"-" json:"-" bson:"-"`
	Id         string   `bson:"_id" json:"id,omitempty" description:"用户ID"`                              //用户ID
	UserName   string   `bson:"UserName" json:"username,omitempty" binding:"required" description:"登陆名"` // 登陆名
	Password   string   `bson:"Password" json:"password,omitempty" description:"登陆密码"`                   // 登陆密码
	Name       string   `bson:"Name" json:"realName,omitempty" description:"用户姓名"`                       // 用户姓名
	CreateTime int64    `bson:"CreateTime" json:"createTime,omitempty" description:"创建时间"`               // 创建时间
	Salt       string   `bson:"Salt" json:"salt,omitempty" description:"密码加密"`                           // 密码加密
	Ids        []string `orm:"-" bson:"-" json:"ids,omitempty"`                                          // Ids
	Roles      []string `bson:"Roles" json:"roles,omitempty" description:"用户角色"`                         // 用户角色
}

func (m *User) One() (errC *resp.Response, err error) {
	err = o.Read(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *User) All() (ls *List, errC *resp.Response, err error) {
	ls = new(List)
	ss := make([]*User, 0)
	qs := o.QueryTable(&User{})
	if m.UserName != "" {
		qs = qs.Filter("UserName__regex", m.UserName)
	}
	if m.Name != "" {
		qs = qs.Filter("Name__regex", m.Name)
	}
	cnt, err := qs.Count()
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	err = qs.OrderBy("Id").Limit(m.PageSize, (m.Page-1)*m.PageSize).All(&ss)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	ls.Lists = ss
	ls.Total = cnt
	return
}

func (m *User) Add() (errC *resp.Response, err error) {
	exist := o.QueryTable(&User{}).Filter("UserName", m.UserName).Exist()
	if exist {
		err = ErrUserExist
		errC = resp.ErrDupRecord
		errC.MoreInfo = err.Error()
		return
	}
	salt, err := auth.GenerateSalt()
	if err != nil {
		errC = resp.Err500
		errC.MoreInfo = err.Error()
		return
	}
	m.Password, err = auth.GeneratePassHash(m.Password, salt)
	if err != nil {
		errC = resp.Err500
		errC.MoreInfo = err.Error()
		return
	}
	m.CreateTime = time.Now().Unix()
	m.Salt = salt

	created, _, err := o.ReadOrCreate(m, "UserName")
	if err != nil {
		errC = resp.ErrDbInsert
		errC.MoreInfo = err.Error()
		return
	}

	if !created {
		err = fmt.Errorf("用户名已存在！")
		errC = resp.ErrDupRecord
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *User) Delete() (errC *resp.Response, err error) {
	_, err = o.Delete(m)
	if err != nil {
		errC = resp.ErrDbDelete
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *User) Update() (errC *resp.Response, err error) {
	exist := o.QueryTable(&User{}).Filter("_id__ne", m.Id).Filter("UserName", m.UserName).Exist()
	if exist {
		err = fmt.Errorf("登录名称重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "登录名称重复!"
		return
	}
	_, err = o.Update(m, "UserName", "Name", "Roles")
	if err != nil {
		errC = resp.ErrDbDelete
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *User) UpdatePassword() (errC *resp.Response, err error) {
	pass := m.Password
	err = o.Read(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	m.Password, err = auth.GeneratePassHash(pass, m.Salt)
	if err != nil {
		errC = resp.Err500
		errC.MoreInfo = err.Error()
		return
	}
	_, err = o.Update(m, "Password")
	if err != nil {
		errC = resp.ErrDbDelete
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *User) Login() (token *Token, errC *resp.Response, err error) {
	pass := m.Password
	err = o.Read(m, "Username")
	if err != nil {
		errC = resp.ErrNoUser
		errC.MoreInfo = err.Error()
		return
	}
	hash, err := auth.GeneratePassHash(pass, m.Salt)
	if err != nil || m.Password != hash {
		err = errors.New("Login Password Error")
		errC = resp.ErrUserPass
		errC.MoreInfo = err.Error()
		return
	}

	token = NewTokenForUser(m.UserName, m.Password)

	err = token.Add()
	if err != nil {
		errC = resp.ErrUserLogin
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *User) Logout(token string) (errC *resp.Response, err error) {
	t := &Token{
		Token: token,
	}
	err = t.Delete()
	if err != nil {
		errC = resp.ErrDbDelete
		errC.MoreInfo = err.Error()
	}
	return
}

func (m *User) GetUserInfo(token string) (errC *resp.Response, err error) {
	t := &Token{
		Token: token,
	}
	err = t.Find()
	if err != nil {
		errC = resp.ErrTokenInValid
		errC.MoreInfo = err.Error()
		return
	}

	username, err := auth.Token_auth(t.Token, t.Secret)
	if err != nil {
		errC = resp.ErrTokenInValid
		errC.MoreInfo = err.Error()
		return
	}
	m.UserName = username
	err = o.Read(m, "UserName")
	if err != nil {
		errC = resp.ErrNoUser
		errC.MoreInfo = err.Error()
		return
	}
	return
}
