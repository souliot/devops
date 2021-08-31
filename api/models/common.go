package models

import (
	"github.com/go-resty/resty/v2"
	"github.com/souliot/siot-orm/orm"
)

var (
	o           orm.Ormer
	httpCli     = resty.New()
	PromAddress string
)

type PageQuery struct {
	Page     int `json:"-"`
	PageSize int `json:"-"`
}

var DefaultPageQuery = &PageQuery{
	PageSize: 0,
	Page:     1,
}

type List struct {
	Total int64       `json:"total"`
	Lists interface{} `json:"lists"`
}

func InitModels() {
	orm.RegisterModel(new(Export))
	orm.RegisterModel(new(Environment))
	orm.RegisterModel(new(PromJob))
	o = orm.NewOrm()
}
