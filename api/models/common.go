package models

import (
	"github.com/souliot/siot-orm/orm"
)

var (
	o orm.Ormer
)

type PageQuery struct {
	Page     int `json:"-"`
	PageSize int `json:"-"`
}

type List struct {
	Total int64       `json:"total"`
	Lists interface{} `json:"lists"`
}

func InitModels() {
	orm.RegisterModel(new(Export))
	orm.RegisterModel(new(Environment))
	o = orm.NewOrm()
}
