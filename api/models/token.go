package models

import (
	"devops/pkg/auth"
	"time"
)

type Token struct {
	Id         string `bson:"_id" json:"id,omitempty"`                //Id
	Token      string `bson:"Token" json:"token,omitempty"`           // token 内容
	Secret     string `bson:"Secret" json:"secret,omitempty"`         // Secret
	CreateTime int64  `bson:"CreateTime" json:"createTime,omitempty"` //创建时间
}

func NewTokenForApp(appid string, secret string) (t *Token) {
	token, _ := auth.Create_token(appid, secret, int64(TokenExp))
	t = &Token{
		Token:      token,
		Secret:     secret,
		CreateTime: time.Now().Unix(),
	}
	return
}

func NewTokenForUser(username string, password string) (t *Token) {
	now := time.Now()
	secret := auth.To_md5(password + now.Format(FormatDateTime))

	token, _ := auth.Create_token(username, secret, int64(TokenExp))
	t = &Token{
		Token:      token,
		Secret:     secret,
		CreateTime: now.Unix(),
	}
	return
}

func (m *Token) Add() (err error) {
	_, _, err = o.ReadOrCreate(m, "Token")
	return
}
func (m *Token) Delete() (err error) {
	_, err = o.Delete(m, "Token")
	return
}

func (m *Token) Find() (err error) {
	err = o.Read(m, "Token")
	return
}

func (m *Token) Valid() (valid bool) {
	err := m.Find()
	if err != nil {
		return false
	}

	_, err = auth.Token_auth(m.Token, m.Secret)
	if err != nil {
		return false
	}

	return true
}
