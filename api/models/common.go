package models

import (
	"github.com/souliot/siot-orm/orm"
)

var (
	o orm.Ormer
)

func InitModels() {
	orm.RegisterModel(new(Export))
	o = orm.NewOrm()
}
