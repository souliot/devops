package models

import (
	"devops/pkg/resp"
)

type Export struct {
	Id         string   `bson:"_id"`
	Type       string   `bson:"Type"`
	Address    string   `bson:"Address"`
	CreateTime int64    `bson:"CreateTime"`
	Targets    []string `orm:"-" json:"targets"`
}

func (m *Export) Node() (exs []*Export, errC *resp.Response, err error) {
	exs = make([]*Export, 0)
	ex := &Export{
		Targets: []string{"192.168.0.252:8080"},
	}
	exs = append(exs, ex)
	return
}
