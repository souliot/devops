package config

import (
	"fmt"
	"strings"

	logs "github.com/souliot/siot-log"
)

func InitLog(cfg *ServerCfg) {
	appname := cfg.AppName
	addr := fmt.Sprintf("%s:%d", cfg.LocalIP, cfg.HttpPort)
	logs.SetLogFuncCall(true)
	logs.SetLevel(cfg.LogLevel)
	logs.EnableFullFilePath(false)
	logs.WithPrefix(appname)
	logs.WithPrefix(addr)
	filepath := strings.TrimRight(cfg.LogPath, "/") + "/" + appname + ".log"
	logs.SetLogger("file", `{"filename":"`+filepath+`","daily":true,"maxdays":10,"color":false}`)
	logs.SetLogger("console")
}
