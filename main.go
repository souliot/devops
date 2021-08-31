package main

import (
	"devops/api"
	"devops/api/config"
	_ "devops/docs"
	"os"
	"os/signal"
	"syscall"

	logs "github.com/souliot/siot-log"
)

var (
	appname = "devops"
	version = "5.2.0.0"
)

// @title  DevOps 开发文档
// @version v1.0.0
// @description  Golang api of demo
// @termsOfService github.com/souliot

// @contact.name API Support
// @contact.url github.com/souliot
// @contact.email leizhou.lin@watrix.ai
// @BasePath /v1
func main() {
	srv := api.NewServer(config.WithAppName(appname), config.WithVersion(version))
	logs.SetLogFuncCall(true)
	srv.Start()

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	_ = <-chSig
}
