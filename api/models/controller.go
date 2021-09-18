package models

import (
	"devops/pkg/resp"
	"fmt"
	"public/libs_go/gateway/master"
)

type Controller struct {
	Env string
	Id  string
}

func (m *Controller) Controller() (errC *resp.Response, err error) {
	msi, loaded := DefaultService.watchCache.Load(m.Env)
	if !loaded {
		err = fmt.Errorf("环境监控错误！")
		errC = resp.ErrEtcdGet
		errC.MoreInfo = err.Error()
		return
	}
	ms, ok := msi.(*master.Master)
	if !ok {
		err = fmt.Errorf("环境监控错误！")
		errC = resp.ErrEtcdGet
		errC.MoreInfo = err.Error()
		return
	}
	cv := &master.ControllerValue{
		Typ: master.ControllerRestart,
	}
	ms.PutController(m.Id, cv)
	return
}
