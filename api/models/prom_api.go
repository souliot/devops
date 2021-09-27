package models

import (
	"fmt"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var (
	promApi v1.API
)

type PromApi struct{}

func initPromApi() (err error) {
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("http://%s", PromAddress),
	})
	if err != nil {
		return
	}
	promApi = v1.NewAPI(client)
	return
}
