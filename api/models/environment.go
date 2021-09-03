package models

import (
	"devops/pkg/resp"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/souliot/siot-orm/orm"

	"github.com/souliot/gateway/master"
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
	cond := orm.NewCondition()
	cond = cond.And("Name", m.Name).Or("EtcdEndpoints", m.EtcdEndpoints)
	exist := o.QueryTable(&Environment{}).SetCond(cond).Exist()
	if exist {
		err = fmt.Errorf("环境名称及地址重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "环境名称及地址重复!"
		return
	}
	m.CreateTime = time.Now().Unix()
	_, err = o.Insert(m)
	if err != nil {
		errC = resp.ErrDbInsert
		errC.MoreInfo = err.Error()
		return
	}
	go m.Watch(DefaultService)
	return
}

func (m *Environment) All() (ls *List, errC *resp.Response, err error) {
	if m.PageQuery == nil {
		m.PageQuery = DefaultPageQuery
	}
	ls = new(List)
	res := make([]*Environment, 0)
	qs := o.QueryTable(&Environment{})
	if m.Name != "" {
		qs = qs.Filter("Name__regex", m.Name)
	}

	cnt, err := qs.Count()
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
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
	DefaultService.StopEnv(m.Name)
	return
}

func (m *Environment) Update() (errC *resp.Response, err error) {
	m_old := new(Environment)
	m_old.Id = m.Id
	err = o.Read(m_old)
	if err != nil {
		errC = resp.ErrNoRecord
		errC.MoreInfo = err.Error()
		return
	}
	cond := orm.NewCondition()
	cond1 := orm.NewCondition()
	cond1 = cond1.And("Name", m.Name).Or("EtcdEndpoints", m.EtcdEndpoints)
	cond = cond.And("_id__ne", m.Id).AndCond(cond1)
	exist := o.QueryTable(&Environment{}).SetCond(cond).Exist()
	if exist {
		err = fmt.Errorf("环境名称及地址重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "环境名称及地址重复!"
		return
	}

	_, err = o.Update(m, "Name", "EtcdEndpoints", "Desc")
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	addrs_old := sort.StringSlice(m_old.EtcdEndpoints)
	addrs := sort.StringSlice(m.EtcdEndpoints)
	sort.Sort(addrs_old)
	sort.Sort(addrs)

	if !reflect.DeepEqual(addrs_old, addrs) {
		DefaultService.StopEnv(m.Name)
		go m.Watch(DefaultService)
	}
	return
}

func (m *Environment) Watch(ser *Service) (err error) {
	if len(m.EtcdEndpoints) <= 0 {
		return
	}
	op := &master.ServiceOption{}
	ms, err := master.OnWatchService(m.EtcdEndpoints, op, 10*time.Second)
	if err != nil {
		return
	}
	go func() {
		for {
			select {
			case <-ms.IsUpdate:
			}
		}
	}()
	if ser.watchCache != nil {
		ser.watchCache.Store(m.Name, ms)
	}
	return
}
