package models

import (
	e "public/entities"
	"public/models/sconf"
	"strconv"
)

type Setting struct{}

func GetTypSettingStruct(typ string) (s interface{}) {
	switch typ {
	case strconv.Itoa(e.AccountTypeAlignment):
		return new(sconf.SearchNodeSerSetting)
	case strconv.Itoa(e.AccountTypeBusiness):
		return new(sconf.BusinessServerSerSetting)
	case strconv.Itoa(e.AccountTypeAlarm):
	case strconv.Itoa(e.AccountTypeCom):
	case strconv.Itoa(e.AccountTypeMaster):
		return new(sconf.MasterNodeSerSetting)
	case strconv.Itoa(e.AccountTypeFS):
		return new(sconf.FsSerSetting)
	case strconv.Itoa(e.AccountTypeCollect):
	case strconv.Itoa(e.AccountTypeSupport):
	case strconv.Itoa(e.AccountTypeMonitor):
	case strconv.Itoa(e.AccountTypeDistribution):
		return new(sconf.AshardsSerSetting)
	case strconv.Itoa(e.AccountTypeRtmp):
	case strconv.Itoa(e.AccountTypeCaseVideo):
	case strconv.Itoa(e.AccountTypeDataBackup):
	case strconv.Itoa(e.AccountTypeGB):
	case strconv.Itoa(e.AccountTypeSaveMonitor):
	case strconv.Itoa(e.AccountTypeFeatureStore):
		return new(sconf.FeatureCollectSerSetting)
	case strconv.Itoa(e.AccountTypeSetting):
		return new(sconf.SettingServiceSerSetting)
	case strconv.Itoa(e.AccountTypeVideoScheduler):
		return new(sconf.VideoManagerSerSetting)
	}
	return
}

func GetAppSettingStruct(typ string) (s interface{}) {
	switch typ {
	case strconv.Itoa(e.AccountTypeAlignment):
		return new(sconf.SearchNodeAppSetting)
	case strconv.Itoa(e.AccountTypeBusiness):
		return new(sconf.BusinessServerAppSetting)
	case strconv.Itoa(e.AccountTypeAlarm):
	case strconv.Itoa(e.AccountTypeCom):
	case strconv.Itoa(e.AccountTypeMaster):
		return new(sconf.MasterNodeAppSetting)
	case strconv.Itoa(e.AccountTypeFS):
		return new(sconf.FsAppSetting)
	case strconv.Itoa(e.AccountTypeCollect):
	case strconv.Itoa(e.AccountTypeSupport):
	case strconv.Itoa(e.AccountTypeMonitor):
	case strconv.Itoa(e.AccountTypeDistribution):
	case strconv.Itoa(e.AccountTypeRtmp):
	case strconv.Itoa(e.AccountTypeCaseVideo):
	case strconv.Itoa(e.AccountTypeDataBackup):
	case strconv.Itoa(e.AccountTypeGB):
	case strconv.Itoa(e.AccountTypeSaveMonitor):
	case strconv.Itoa(e.AccountTypeFeatureStore):
		return new(sconf.FeatureCollectAppSetting)
	case strconv.Itoa(e.AccountTypeSetting):
		return new(sconf.SettingServiceAppSetting)
	case strconv.Itoa(e.AccountTypeVideoScheduler):
		return new(sconf.VideoManageAppSetting)
	}
	return
}
