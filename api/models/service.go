package models

import (
	"devops/pkg/resp"
	"public/libs_go/gateway/master"
	"sort"
	"strconv"
	"sync"
)

type ServiceRequest struct {
	Env         string `json:"env"`
	Path        string `json:"path"`
	Typ         string `json:"typ"`
	Id          string `json:"id"`
	OutAddress  string `json:"outAddress"`
	MetricsType string `json:"metricsType"`
}

var DefaultService = &Service{
	watchCache: new(sync.Map),
}

type Service struct {
	watchCache *sync.Map
}

type ServiceMeta master.ServiceMeta

type ServiceResponse struct {
	ServiceMeta
	Env      string
	IsOnline bool
}

func (m *Service) Watch() {
	env := new(Environment)
	ls, _, err := env.All()
	if err != nil {
		return
	}
	for _, v := range ls.Lists.([]*Environment) {
		go v.Watch(m)
	}
}

func (m *Service) All(req *ServiceRequest) (ls *List, errC *resp.Response, err error) {
	op := &master.ServiceOption{
		Path:        req.Path,
		Typ:         req.Typ,
		Id:          req.Id,
		MetricsType: master.MetricsType(req.MetricsType),
	}
	ls = new(List)
	res := make([]*ServiceResponse, 0)
	if req.Env == "" {
		m.watchCache.Range(func(k, v interface{}) bool {
			if ms, ok := v.(*master.Master); ok {
				sms, err := ms.GetNodes(op)
				if err != nil {
					return true
				}
				env := k.(string)
				for _, v := range sms {
					sr := &ServiceResponse{
						ServiceMeta: ServiceMeta(v.Meta),
						IsOnline:    v.Status,
						Env:         env,
					}
					res = append(res, sr)
				}
			}
			return true
		})
		sort.Sort(ServiceResponseList(res))
		ls.Total = int64(len(res))
		ls.Lists = res
		return
	}
	if msi, loaded := m.watchCache.Load(req.Env); loaded {
		if ms, ok := msi.(*master.Master); ok {
			sms, err := ms.GetNodes(op)
			if err != nil {
				errC = resp.ErrEtcdGet
				errC.MoreInfo = err.Error()
				return nil, errC, err
			}
			for _, v := range sms {
				sr := &ServiceResponse{
					ServiceMeta: ServiceMeta(v.Meta),
					IsOnline:    v.Status,
					Env:         req.Env,
				}
				res = append(res, sr)
			}
		}
	}
	sort.Sort(ServiceResponseList(res))
	ls.Total = int64(len(res))
	ls.Lists = res
	return
}

func (m *Service) Online(req *ServiceRequest) (ls *List, errC *resp.Response, err error) {
	op := &master.ServiceOption{
		Path:        req.Path,
		Typ:         req.Typ,
		Id:          req.Id,
		MetricsType: master.MetricsType(req.MetricsType),
	}
	ls = new(List)
	res := make([]*ServiceResponse, 0)
	if req.Env == "" {
		m.watchCache.Range(func(k, v interface{}) bool {
			if ms, ok := v.(*master.Master); ok {
				sms, err := ms.GetServicesOnline(op)
				if err != nil {
					return true
				}
				env := k.(string)
				for _, v := range sms {
					sr := &ServiceResponse{
						ServiceMeta: ServiceMeta(*v),
						IsOnline:    true,
						Env:         env,
					}
					res = append(res, sr)
				}
			}
			return true
		})
		return
	}
	if msi, loaded := m.watchCache.Load(req.Env); loaded {
		if ms, ok := msi.(*master.Master); ok {
			sms, err := ms.GetServicesOnline(op)
			if err != nil {
				errC = resp.ErrEtcdGet
				errC.MoreInfo = err.Error()
				return nil, errC, err
			}
			for _, v := range sms {
				sr := &ServiceResponse{
					ServiceMeta: ServiceMeta(*v),
					IsOnline:    true,
					Env:         req.Env,
				}
				res = append(res, sr)
			}
		}
	}
	ls.Total = int64(len(res))
	ls.Lists = res
	return
}

func (m *Service) Stop() {
	m.watchCache.Range(func(k, v interface{}) bool {
		if ms, ok := v.(*master.Master); ok {
			ms.Stop()
		}
		return true
	})
}

func (m *Service) Delete() (errC *resp.Response, err error) {
	m.watchCache.Range(func(k, v interface{}) bool {
		if ms, ok := v.(*master.Master); ok {
			ms.Stop()
		}
		return true
	})
	return
}

func (m *Service) DeleteNode(req *ServiceRequest) (errC *resp.Response, err error) {
	if msi, loaded := m.watchCache.Load(req.Env); loaded {
		if ms, ok := msi.(*master.Master); ok {
			ms.DeleteNode(req.Path, req.Typ, req.Id)
		}
	}
	return
}

func (m *Service) SetOutAddress(req *ServiceRequest) (errC *resp.Response, err error) {
	if msi, loaded := m.watchCache.Load(req.Env); loaded {
		if ms, ok := msi.(*master.Master); ok {
			ms.SetOutAddress(req.Id, req.OutAddress)
		}
	}
	return
}

func (m *Service) StopEnv(name string) {
	if msi, loaded := m.watchCache.LoadAndDelete(name); loaded {
		if ms, ok := msi.(*master.Master); ok {
			go ms.Stop()
		}
	}
}

func (m *Service) GetExport(env, typ string) (exps []string) {
	op := &master.ServiceOption{
		MetricsType: master.MetricsType(typ),
	}
	exps = make([]string, 0)
	if env == "" {
		m.watchCache.Range(func(k, v interface{}) bool {
			if ms, ok := v.(*master.Master); ok {
				sms, err := ms.GetServicesOnline(op)
				if err != nil {
					return true
				}
				for _, sm := range sms {
					addr := sm.OutAddress
					if addr == "" {
						addr = sm.Address
					}
					exps = append(exps, addr)
				}
			}
			return true
		})
		return
	}
	if msi, loaded := m.watchCache.Load(env); loaded {
		if ms, ok := msi.(*master.Master); ok {
			sms, err := ms.GetServicesOnline(op)
			if err != nil {
				return
			}
			for _, sm := range sms {
				addr := sm.OutAddress
				if addr == "" {
					addr = sm.Address
				}
				exps = append(exps, addr)
			}
		}
	}
	return
}

type ServiceResponseList []*ServiceResponse

func (m ServiceResponseList) Len() int {
	return len(m)
}

func (m ServiceResponseList) Less(i, j int) bool {
	// 环境
	if m[i].Env < m[j].Env {
		return true
	}
	if m[i].Env > m[j].Env {
		return false
	}

	// 服务类型

	typA, errA := strconv.Atoi(m[i].Typ)
	typB, errB := strconv.Atoi(m[j].Typ)
	if errA != nil && errB != nil {
		if m[i].Typ < m[j].Typ {
			return true
		}
		if m[i].Typ > m[j].Typ {
			return false
		}
	}
	if errA != nil {
		return true
	}
	if errB != nil {
		return false
	}

	if typA < typB {
		return true
	}
	if typA > typB {
		return false
	}

	// 集群
	if m[i].Path < m[j].Path {
		return true
	}
	if m[i].Path > m[j].Path {
		return false
	}

	if m[i].ServiceMeta.Id <= m[j].ServiceMeta.Id {
		return true
	}
	return false
}

func (m ServiceResponseList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
	return
}
