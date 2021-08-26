package main

import (
	"devops/api"
	"os"
	"os/signal"
	"syscall"

	logs "github.com/souliot/siot-log"
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
	srv := api.NewServer()
	logs.SetLogFuncCall(true)
	srv.Start()

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	_ = <-chSig
}
