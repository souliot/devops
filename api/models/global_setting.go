package models

import (
	"devops/pkg/resp"
	"encoding/json"
	"fmt"
	e "public/entities"
	"public/libs_go/gateway/master"
)

const GLOBAL_SETTING_NAME = "AppSetting"

type GlobalSetting struct {
	Env string
}

func (m *GlobalSetting) GetGlobalSetting() (res *e.ServerSetting, errC *resp.Response, err error) {
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

	data, err := ms.GetTypSettingRaw(GLOBAL_SETTING_NAME)
	if err != nil {
		errC = resp.ErrEtcdGet
		errC.MoreInfo = "配置信息不存在！"
		return
	}
	res = new(e.ServerSetting)
	err = json.Unmarshal(data, res)
	if err != nil {
		errC = resp.ErrTransferData
		errC.MoreInfo = err.Error()
		return nil, errC, err
	}
	return
}

func (m *GlobalSetting) SetGlobalSetting(data []byte) (errC *resp.Response, err error) {
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
	err = ms.PutTypSettingRaw(GLOBAL_SETTING_NAME, data)
	if err != nil {
		errC = resp.ErrEtcdPut
		errC.MoreInfo = err.Error()
	}
	return
}
