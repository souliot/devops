package models

import (
	"devops/pkg/resp"
	"encoding/json"
	"fmt"
	"public/libs_go/gateway/master"
)

type AppSetting struct {
	Env string
	Typ string
	Id  string
}

func (m *AppSetting) GetAppSetting() (res *master.AppSetting, errC *resp.Response, err error) {
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
	ss := &master.AppSetting{
		Data: GetAppSettingStruct(m.Typ),
	}
	ser, err := ms.GetSrvSetting(m.Id, ss)
	if err != nil || ser == nil || ser.AppSetting == nil {
		errC = resp.ErrEtcdGet
		errC.MoreInfo = "配置信息不存在！"
		return
	}
	res = ser.AppSetting
	return
}

func (m *AppSetting) SetAppSetting(data []byte) (errC *resp.Response, err error) {
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
	ss := &master.AppSetting{
		Data: GetAppSettingStruct(m.Typ),
	}
	err = json.Unmarshal(data, ss)
	if err != nil {
		errC = resp.ErrTransferData
		errC.MoreInfo = err.Error()
		return
	}
	err = ms.PutAppSetting(m.Id, ss)
	if err != nil {
		errC = resp.ErrEtcdPut
		errC.MoreInfo = err.Error()
	}
	return
}
