package api

import (
	"devops/api/config"
	"devops/api/models"
	"devops/api/routes"
	"devops/pkg/db"
	"fmt"

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
	models.DefaultMetrics.Init()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	routes.InitRouter(r)
	go r.Run(fmt.Sprintf(":%d", s.cfg.HttpPort))
	logs.Info("API Server 启动成功，端口：%d", s.cfg.HttpPort)
}

func (s *Server) SaveConfig() {
	if s.cfg != nil {
		s.cfg.SaveConfigFile()
	}
}
