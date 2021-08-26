package db

import (
	"strings"

	"github.com/souliot/siot-orm/orm"

	logs "github.com/souliot/siot-log"
)

var (
	mongodb = "default"
)

type MongoSetting struct {
	Hosts    []string `json:"Hosts"`
	Username string   `json:"Username"`
	Password string   `json:"Password"`
	DBName   string   `json:"DbName"`
}

func InitMongo(s *MongoSetting) {
	if orm.HasDefaultDataBase() {
		mongodb = "mongodb"
	}
	mu := ""
	if s.Username != "" {
		mu = s.Username + ":" + s.Password + "@"
	}
	mongo_address := "mongodb://" + mu + strings.Join(s.Hosts, ",") + "/" + s.DBName + "?authSource=admin"

	orm.RegisterDriver("mongo", orm.DRMongo)
	err := orm.RegisterDataBase(mongodb, "mongo", mongo_address, true)
	if err != nil {
		logs.Error("初始化mongodb错误：", err)
		return
	}
	logs.Info("Mongo 数据库初始化成功：", mongo_address)
}
