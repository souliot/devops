package models

import (
	"devops/pkg/resp"
	"time"
)

type Export struct {
	*PageQuery `orm:"-" json:"-" bson:"-"`
	Id         string   `bson:"_id" json:"id,omitempty"`                             // Id
	Type       string   `bson:"Type" binding:"required" json:"type,omitempty"`       // 节点类型
	Address    string   `bson:"Address" binding:"required" json:"address,omitempty"` // 节点地址
	CreateTime int64    `bson:"CreateTime" json:"createTime,omitempty"`              // 创建时间
	Targets    []string `orm:"-" bson:"-" json:"targets,omitempty"`                  // 节点地址列表
}

func (m *Export) Add() (errC *resp.Response, err error) {
	m.CreateTime = time.Now().Unix()
	o.ReadOrCreate(m, "Type", "Address")
	if err != nil {
		errC = resp.ErrDbInsert
		errC.MoreInfo = err.Error()
	}
	return
}

func (m *Export) All() (res []*Export, errC *resp.Response, err error) {
	res = make([]*Export, 0)
	qs := o.QueryTable(&Export{})
	if m.Type != "" {
		qs = qs.Filter("Type", m.Type)
	}
	if m.Address != "" {
		qs = qs.Filter("Address__regex", m.Address)
	}

	err = qs.Limit(m.PageSize, (m.Page-1)*m.PageSize).OrderBy("-CreateTime").All(&res)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Export) One() (errC *resp.Response, err error) {
	err = o.Read(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Export) Delete() (errC *resp.Response, err error) {
	_, err = o.Delete(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Export) Update() (errC *resp.Response, err error) {
	_, err = o.Update(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Export) Node() (ex *Export, errC *resp.Response, err error) {
	exs := make([]*Export, 0)
	err = o.QueryTable("Export").Filter("Type", m.Type).All(&exs)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	targets := make([]string, 0)
	for _, v := range exs {
		targets = append(targets, v.Address)
	}
	ex = &Export{
		Targets: targets,
	}
	return
}
