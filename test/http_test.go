package test

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

var ts http.RoundTripper = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:        20000,
	MaxIdleConnsPerHost: 100,
	IdleConnTimeout:     30 * time.Second,
}

func TestHttpTimeout(t *testing.T) {
	start := time.Now()

	url := "http://192.168.0.35:1231"
	client := &http.Client{Transport: ts}
	_, err := client.Post(url, "", nil)
	if err != nil {
		t.Fatal(err, time.Since(start))
	}
	t.Log(time.Since(start))
}

func TestDefer(t *testing.T) {
	defer fmt.Println("aaaa")
	defer fmt.Println("bbbb")
	defer fmt.Println("cccc")
	fmt.Println("dddddddd")
}
