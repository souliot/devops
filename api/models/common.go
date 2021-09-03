package models

import (
	"github.com/go-resty/resty/v2"
	"github.com/souliot/siot-orm/orm"
)

const (
	FormatTime     = "15:04:05"
	FormatDate     = "2006-01-02"
	FormatDateTime = "2006-01-02 15:04:05"
)

var (
	o           orm.Ormer
	httpCli     = resty.New()
	PromAddress string
	TokenExp    int
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
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Token))
	orm.RegisterModel(new(AppUser))
	o = orm.NewOrm()
	InitData()
}

func InitData() (err error) {
	// default user
	u := &User{
		UserName: "llz",
		Password: "1",
		Name:     "Souliot",
	}
	_, err = u.Add()
	if err != nil && err.Error() != ErrUserExist.Error() {
		return
	}

	// default appuser
	appuser := new(AppUser)
	err = appuser.Add(u)
	if err != nil {
		return
	}

	return
}
