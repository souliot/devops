package models

import (
	"devops/pkg/auth"
	"devops/pkg/resp"
	"time"
)

type AppUser struct {
	Id         string   `bson:"_id" json:"id"`                          // Id
	Appid      string   `bson:"Appid"  json:"appid,omitempty"`          // 应用Id
	Secret     string   `bson:"Secret"  json:"secret,omitempty"`        // Secret
	CreateTime int64    `bson:"CreateTime" json:"createTime,omitempty"` // 创建时间
	Salt       string   `bson:"Salt"  json:"salt,omitempty"`            // Salt
	Roles      []string `bson:"Roles" json:"roles,omitempty"`           // 应用角色
	UserId     string   `bson:"UserId"  json:"userId,omitempty"`        // 用户Id
	Access     []string `orm:"-" bson:"Access" json:"access,omitempty"` // 用户角色
}

func (m *AppUser) Add(a *User) (err error) {
	salt, err := auth.GenerateSalt()
	if err != nil {
		return
	}

	pass, err := auth.GeneratePassHash(a.Password, salt)
	if err != nil {
		return
	}
	now := time.Now()

	m.Appid = auth.To_md5(a.UserName + now.Format(FormatDateTime))[0:16]
	m.Secret = auth.To_md5(pass + now.Format(FormatDateTime))
	m.Salt = salt
	m.Roles = []string{"admin"}
	m.UserId = a.Id
	m.CreateTime = now.Unix()

	_, _, err = o.ReadOrCreate(m, "UserId")
	return
}

func (m *AppUser) Delete() (err error) {
	_, err = o.Delete(m)

	return
}

func (m *AppUser) DeleteByUser() (err error) {
	qs := o.QueryTable(&AppUser{})
	_, err = qs.Filter("UserId", m.UserId).Delete()
	return
}

func (m *AppUser) GetByUser() (errC *resp.Response, err error) {
	err = o.Read(m, "UserId")
	if err != nil {
		errC = resp.ErrNoRecord
	}
	return
}

func (m *AppUser) Valid() (valid bool) {
	err := o.Read(m, "Appid", "Secret")

	if err != nil {
		return false
	}
	return true
}
