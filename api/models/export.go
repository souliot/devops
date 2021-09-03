package models

import (
	"devops/pkg/resp"
	"fmt"
	"time"

	"github.com/souliot/siot-orm/orm"
)

type Export struct {
	*PageQuery `orm:"-" json:"-" bson:"-"`
	Id         string   `bson:"_id" json:"id,omitempty"`                             // Id
	Env        string   `bson:"Env" json:"env,omitempty"`                            // 环境名称
	Type       string   `bson:"Type" binding:"required" json:"type,omitempty"`       // 节点类型
	Address    string   `bson:"Address" binding:"required" json:"address,omitempty"` // 节点地址
	CreateTime int64    `bson:"CreateTime" json:"createTime,omitempty"`              // 创建时间
	Targets    []string `orm:"-" bson:"-" json:"targets,omitempty"`                  // 节点地址列表
}

func (m *Export) Add() (errC *resp.Response, err error) {
	cond := orm.NewCondition()
	cond = cond.And("Type", m.Type).And("Address", m.Address)
	exist := o.QueryTable(&Export{}).SetCond(cond).Exist()
	if exist {
		err = fmt.Errorf("节点类型及地址重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "节点类型及地址重复!"
		return
	}
	m.CreateTime = time.Now().Unix()
	_, err = o.Insert(m)
	if err != nil {
		errC = resp.ErrDbInsert
		errC.MoreInfo = err.Error()
	}
	return
}

func (m *Export) All() (ls *List, errC *resp.Response, err error) {
	if m.PageQuery == nil {
		m.PageQuery = DefaultPageQuery
	}
	ls = new(List)
	res := make([]*Export, 0)
	qs := o.QueryTable(&Export{})
	if m.Type != "" {
		qs = qs.Filter("Type", m.Type)
	}
	if m.Address != "" {
		qs = qs.Filter("Address__regex", m.Address)
	}

	cnt, err := qs.Count()
	if err != nil {
		errC = resp.ErrDbRead
		return
	}

	err = qs.Limit(m.PageSize, (m.Page-1)*m.PageSize).OrderBy("CreateTime").All(&res)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	ls.Lists = res
	ls.Total = cnt
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
	cond := orm.NewCondition()
	cond = cond.And("_id__ne", m.Id).And("Type", m.Type).And("Address", m.Address)
	exist := o.QueryTable(&Export{}).SetCond(cond).Exist()
	if exist {
		err = fmt.Errorf("节点类型及地址重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "节点类型及地址重复!"
		return
	}
	_, err = o.Update(m, "Env", "Type", "Address")
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Export) Node() (res []*Export, errC *resp.Response, err error) {
	exs := make([]*Export, 0)
	qs := o.QueryTable(&Export{})
	if m.Env != "" {
		qs = qs.Filter("Env", m.Env)
	}
	err = qs.Filter("Type", m.Type).All(&exs)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	targets := make([]string, 0)
	tgs_cache := make(map[string]struct{})
	for _, v := range exs {
		tgs_cache[v.Address] = struct{}{}
	}
	eps := DefaultService.GetExport(m.Env, m.Type)

	for _, v := range eps {
		tgs_cache[v] = struct{}{}
	}

	for v, _ := range tgs_cache {
		targets = append(targets, v)
	}

	ex := &Export{
		Targets: targets,
	}
	if len(ex.Targets) <= 0 {
		err = fmt.Errorf("该类型未配置监控目标！")
		errC = resp.ErrNoRecord
		errC.MoreInfo = "该类型未配置监控目标！"
		return
	}
	res = append(res, ex)
	return
}
