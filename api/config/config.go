package config

import (
	"bytes"
	"encoding/json"
	"fmt"

	logs "github.com/souliot/siot-log"
	"github.com/spf13/viper"
)

var (
	errConfigNotInit = fmt.Errorf("config have not init")
)

type ServerCfg struct {
	HttpPort   int      `mapstructure:"httpport"`
	DBHost     []string `mapstructure:"dbhost"`
	DBName     string   `mapstructure:"dbname"`
	DBUser     string   `mapstructure:"dbuser"`
	DBPassword string   `mapstructure:"dbpassword"`
}

var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.SetConfigType("yaml")
	b, _ := json.Marshal(DefaultServerCfg)
	defaultConfig := bytes.NewReader(b)
	Config.ReadConfig(defaultConfig)
	Config.SetConfigFile("config.yaml")
	err := Config.ReadInConfig()
	if err != nil {
		logs.Info("Using default config")
	} else {
		Config.MergeInConfig()
	}

	Config.Unmarshal(DefaultServerCfg)
}

type Option func(*ServerCfg)

var DefaultServerCfg = &ServerCfg{
	HttpPort:   8080,
	DBHost:     []string{"localhost:27017"},
	DBName:     "",
	DBUser:     "",
	DBPassword: "",
}

func (c *ServerCfg) Apply(opts []Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *ServerCfg) SaveConfigFile() (err error) {
	if Config == nil {
		return errConfigNotInit
	}
	err = Config.WriteConfigAs(Config.ConfigFileUsed())
	return
}

func WithHttpPort(p int) Option {
	return func(c *ServerCfg) {
		c.HttpPort = p
	}
}

func WithDBHost(p []string) Option {
	return func(c *ServerCfg) {
		c.DBHost = p
	}
}

func WithDBName(p string) Option {
	return func(c *ServerCfg) {
		c.DBName = p
	}
}

func WithDBUser(p string) Option {
	return func(c *ServerCfg) {
		c.DBUser = p
	}
}

func WithDBPassword(p string) Option {
	return func(c *ServerCfg) {
		c.DBPassword = p
	}
}
