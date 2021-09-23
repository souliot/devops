package models

import (
	"devops/pkg/resp"
	"fmt"
	"time"

	"github.com/souliot/siot-orm/orm"
)

type PromJob struct {
	*PageQuery  `orm:"-" json:"-" bson:"-"`
	Id          string   `bson:"_id" json:"id,omitempty"`                                     // Id
	JobName     string   `bson:"JobName" binding:"required" json:"jobName,omitempty"`         // 采集任务名称
	MetricsPath string   `bson:"MetricsPath" binding:"required" json:"metricsPath,omitempty"` // Metrics Path
	Scheme      string   `bson:"Scheme" binding:"required" json:"scheme,omitempty"`           // 协议 Scheme
	ConfigsType string   `bson:"ConfigsType" binding:"required" json:"configsType,omitempty"` // 配置类型：static || http
	Targets     []string `bson:"Targets" json:"targets,omitempty"`                            // static 采集目标
	Url         string   `bson:"Url" json:"url,omitempty"`                                    // http 请求地址
	CreateTime  int64    `bson:"CreateTime" json:"createTime,omitempty"`                      // 创建时间
}

func (m *PromJob) Add() (errC *resp.Response, err error) {
	cond := orm.NewCondition()
	cond = cond.And("JobName", m.JobName)
	if len(m.Targets) > 0 {
		cond = cond.Or("Targets", m.Targets)
	}
	if m.Url != "" {
		cond = cond.Or("Url", m.Url)
	}
	exist := o.QueryTable(&PromJob{}).SetCond(cond).Exist()
	if exist {
		err = fmt.Errorf("任务名称及地址重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "任务名称及地址重复!"
		return
	}
	switch m.ConfigsType {
	case "http":
		m.Targets = make([]string, 0)
	case "static":
		m.Url = ""
	default:
	}
	m.CreateTime = time.Now().Unix()
	_, err = o.Insert(m)
	if err != nil {
		errC = resp.ErrDbInsert
		errC.MoreInfo = err.Error()
		return
	}
	go AutoReload()
	return
}

func (m *PromJob) All() (ls *List, errC *resp.Response, err error) {
	if m.PageQuery == nil {
		m.PageQuery = DefaultPageQuery
	}
	ls = new(List)
	res := make([]*PromJob, 0)
	qs := o.QueryTable(&PromJob{})
	if m.JobName != "" {
		qs = qs.Filter("JobName__regex", m.JobName)
	}

	cnt, err := qs.Count()
	if err != nil {
		errC = resp.ErrDbRead
		return
	}

	err = qs.Limit(m.PageSize, (m.Page-1)*m.PageSize).OrderBy("JobName").All(&res)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}

	ls.Lists = res
	ls.Total = cnt
	return
}

func (m *PromJob) One() (errC *resp.Response, err error) {
	err = o.Read(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *PromJob) Delete() (errC *resp.Response, err error) {
	_, err = o.Delete(m)
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	go AutoReload()
	return
}

func (m *PromJob) Update() (errC *resp.Response, err error) {
	cond := orm.NewCondition()
	cond = cond.And("_id__ne", m.Id).And("JobName", m.JobName)
	if len(m.Targets) > 0 {
		cond = cond.And("Targets", m.Targets)
	}
	if m.Url != "" {
		cond = cond.And("Url", m.Url)
	}
	exist := o.QueryTable(&PromJob{}).SetCond(cond).Exist()
	if exist {
		err = fmt.Errorf("任务名称及地址重复!")
		errC = resp.ErrDupRecord
		errC.MoreInfo = "任务名称及地址重复!"
		return
	}
	switch m.ConfigsType {
	case "http":
		m.Targets = make([]string, 0)
	case "static":
		m.Url = ""
	default:
	}
	_, err = o.Update(m, "JobName", "MetricsPath", "Scheme", "ConfigsType", "Targets", "Url")
	if err != nil {
		errC = resp.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	go AutoReload()
	return
}
