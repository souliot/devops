package main

import (
	_ "common/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	tem,err := host.SensorsTemperatures()
	if err!=nil{
		logs.Error(err)
	}
	logs.Info(tem)
	beego.Run()
}
