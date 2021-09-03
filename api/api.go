package api

import (
	"devops/api/config"
	"devops/api/models"
	"devops/api/routes"
	"devops/pkg/db"
	"devops/pkg/trans"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	logs "github.com/souliot/siot-log"
)

type Server struct {
	cfg *config.ServerCfg
}

func NewServer(ops ...config.Option) (m *Server) {
	config.InitConfig()
	cfg := config.DefaultServerCfg
	cfg.Apply(ops)
	if cfg.LocalIP == "" {
		cfg.LocalIP = config.GetIPStr()
	}
	config.InitLog(cfg)
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() {
	ms := &db.MongoSetting{
		Hosts:    s.cfg.DBHost,
		Username: s.cfg.DBUser,
		Password: s.cfg.DBPassword,
		DBName:   s.cfg.DBName,
	}
	db.InitMongo(ms)
	trans.InitTrans("zh")
	models.DefaultMetrics.Init(s.cfg.AppName)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	routes.InitRouter(r)
	models.InitModels()
	models.DefaultService.Watch()
	paddr := s.cfg.PromAddress
	paddr = strings.TrimPrefix(paddr, "http://")
	paddr = strings.TrimPrefix(paddr, "https://")
	models.PromAddress = paddr
	models.TokenExp = s.cfg.TokenExp
	go r.Run(fmt.Sprintf(":%d", s.cfg.HttpPort))
	logs.Info("API Server 启动成功，端口：%d", s.cfg.HttpPort)
}

func (s *Server) SaveConfig() {
	if s.cfg != nil {
		s.cfg.SaveConfigFile()
	}
}
