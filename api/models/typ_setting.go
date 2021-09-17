package models

import (
	"devops/pkg/resp"
	"fmt"
	"public/libs_go/gateway/master"
)

type TypSetting struct {
	Env string
	Typ string
}

func (m *TypSetting) GetTypSetting() (res *master.SerSetting, errC *resp.Response, err error) {
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
	ss := &master.SerSetting{
		Data: GetTypSettingStruct(m.Typ),
	}
	res, err = ms.GetTypSetting(m.Typ, ss)
	if err != nil {
		errC = resp.ErrEtcdGet
		errC.MoreInfo = "配置信息不存在！"
		return
	}
	return
}

func (m *TypSetting) SetTypSetting(data []byte) (errC *resp.Response, err error) {
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
	err = ms.PutTypSettingRaw(m.Typ, data)
	if err != nil {
		errC = resp.ErrEtcdPut
		errC.MoreInfo = err.Error()
	}
	return
}
