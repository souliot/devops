package models

import (
	"devops/pkg/resp"
	"sync"

	"github.com/souliot/gateway/master"
	logs "github.com/souliot/siot-log"
)

type ServiceRequest struct {
	Env         string
	Path        string
	Typ         string
	Id          string
	MetricsType string
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
	Status bool
}

func (m *Service) Watch() {
	env := new(Environment)
	evs, _, err := env.All()
	if err != nil {
		return
	}
	for _, v := range evs {
		v.Watch(m)
	}
}

func (m *Service) All(req *ServiceRequest) (res []*ServiceResponse, errC *resp.Response, err error) {
	op := &master.ServiceOption{
		Path:        req.Path,
		Typ:         req.Typ,
		Id:          req.Id,
		MetricsType: master.MetricsType(req.MetricsType),
	}
	res = make([]*ServiceResponse, 0)
	if req.Env == "" {
		m.watchCache.Range(func(k, v interface{}) bool {
			if ms, ok := v.(*master.Master); ok {
				sms, err := ms.GetNodes(op)
				if err != nil {
					return true
				}
				for _, v := range sms {
					sr := &ServiceResponse{
						ServiceMeta: ServiceMeta(v.Meta),
						Status:      v.Status,
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
			sms, err := ms.GetNodes(op)
			if err != nil {
				errC = resp.ErrEtcdGet
				errC.MoreInfo = err.Error()
				return res, errC, err
			}
			for _, v := range sms {
				sr := &ServiceResponse{
					ServiceMeta: ServiceMeta(v.Meta),
					Status:      v.Status,
				}
				res = append(res, sr)
			}
		}
	}
	return
}

func (m *Service) Online(req *ServiceRequest) (res []*master.ServiceMeta, errC *resp.Response, err error) {
	op := &master.ServiceOption{
		Path:        req.Path,
		Typ:         req.Typ,
		Id:          req.Id,
		MetricsType: master.MetricsType(req.MetricsType),
	}
	res = make([]*master.ServiceMeta, 0)
	if req.Env == "" {
		m.watchCache.Range(func(k, v interface{}) bool {
			if ms, ok := v.(*master.Master); ok {
				sms, err := ms.GetServicesOnline(op)
				if err != nil {
					return true
				}
				res = append(res, sms...)
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
				return res, errC, err
			}
			res = append(res, sms...)
		}
	}
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

func (m *Service) StopEnv(name string) {
	logs.Info("Stop Env")
	if msi, loaded := m.watchCache.LoadAndDelete(name); loaded {
		if ms, ok := msi.(*master.Master); ok {
			ms.Stop()
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
					exps = append(exps, sm.MetricsAddress)
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
				exps = append(exps, sm.MetricsAddress)
			}
		}
	}
	return
}
