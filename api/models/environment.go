package models

import (
	"devops/pkg/resp"
	"time"
)

type Environment struct {
	*PageQuery    `orm:"-" json:"-" bson:"-"`
	Id            string   `bson:"_id" json:"id,omitempty"`                                         // Id
	Name          string   `bson:"Name" binding:"required" json:"name,omitempty"`                   // 环境名称
	EtcdEndpoints []string `bson:"EtcdEndpoints" binding:"required" json:"etcdEndpoints,omitempty"` // ETCD地址
	Desc          string   `bson:"Desc" json:"desc,omitempty"`                                      // 环境描述
	CreateTime    int64    `bson:"CreateTime" json:"createTime,omitempty"`                          // 创建时间
}

func (m *Environment) Add() (errC *resp.Response, err error) {
	m.CreateTime = time.Now().Unix()
	o.ReadOrCreate(m, "Name", "EtcdEndpoints")
	if err != nil {
		errC = resp.ErrDbInsert
		errC.MoreInfo = err.Error()
	}
	return
}

func (m *Environment) All() (res []*Environment, errC *resp.Response, err error) {
	res = make([]*Environment, 0)
	qs := o.QueryTable(&Environment{})
	if m.Name != "" {
		qs = qs.Filter("Name__regex", m.Name)
	}

	err = qs.Limit(m.PageSize, (m.Page-1)*m.PageSize).OrderBy("-CreateTime").All(&res)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Environment) One() (errC *resp.Response, err error) {
	err = o.Read(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Environment) Delete() (errC *resp.Response, err error) {
	_, err = o.Delete(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Environment) Update() (errC *resp.Response, err error) {
	_, err = o.Update(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}
