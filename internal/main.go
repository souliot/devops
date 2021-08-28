package main

import (
	"net"
	"net/http"
	"time"

	logs "github.com/souliot/siot-log"
)

var ts http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

func main() {
	start := time.Now()

	url := "http://192.168.0.35:1231"
	client := &http.Client{Transport: ts}
	_, err := client.Post(url, "", nil)
	if err != nil {
		logs.Error(err)
	}
	logs.Info(time.Since(start))
}
