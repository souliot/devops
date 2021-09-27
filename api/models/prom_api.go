package models

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

var (
	promApi v1.API
)

func InitPromApi() (err error) {
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("http://%s", PromAddress),
	})
	if err != nil {
		return
	}
	promApi = v1.NewAPI(client)
	return
}

func getMatrix(pql string, ctx context.Context, r v1.Range) (ret []model.SamplePair, err error) {
	ret = make([]model.SamplePair, 0)
	res, _, err := promApi.QueryRange(ctx, pql, r)
	if err != nil {
		return
	}
	if mat, ok := res.(model.Matrix); ok {
		if mat.Len() < 1 {
			return
		}
		return mat[0].Values, nil
	}
	return
}

func getVector(pql string, ctx context.Context) (ret model.Vector, err error) {
	res, _, err := promApi.Query(ctx, pql, time.Now())
	if err != nil {
		return
	}
	if vec, ok := res.(model.Vector); ok {
		ret = vec
		return
	}
	return nil, fmt.Errorf("Result are not vector")
}
