package models

import (
	"devops/pkg/resp"
	"time"
)

type Export struct {
	Id         string   `bson:"_id" json:"id,omitempty"`
	Type       string   `bson:"Type" binding:"required" json:"type,omitempty"`
	Address    string   `bson:"Address" binding:"required" json:"address,omitempty"`
	CreateTime int64    `bson:"CreateTime" json:"createTime,omitempty"`
	Targets    []string `orm:"-" bson:"-" json:"targets,omitempty"`
}

func (m *Export) Add() (errC *resp.Response, err error) {
	m.CreateTime = time.Now().Unix()
	o.ReadOrCreate(m, "Type", "Address")
	if err != nil {
		errC = resp.ErrDbInsert
	}
	return
}

func (m *Export) Node() (exs []*Export, errC *resp.Response, err error) {
	exs = make([]*Export, 0)
	ex := &Export{
		Targets: []string{"192.168.0.252:8080"},
	}
	exs = append(exs, ex)
	return
}
