package rest_api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-wire-example/internal/services/base"
)

type Config struct {
	ServerAddress string `yaml:"server_address"`
}

type Service struct {
	*base.Service

	Config *Config

	httpRouter *gin.Engine
	httpServer *http.Server
}

func (s *Service) ConfigureRoute() {
	s.httpRouter.GET("/restapi/ping", s.Ping())
}

func (s *Service) Start() {
	s.Service.Start()

	go func() {
		s.ConfigureRoute()
		s.Logger.Infof("starting HTTP server at http://%v ...", s.httpServer.Addr)
		s.Logger.Fatal(s.httpServer.ListenAndServe())
	}()
}

func (s *Service) Stop() {

}

func NewService(
	config *Config,
	baseService *base.Service,
) (*Service, error) {
	httpRouter := gin.New()

	httpAddr := ":8080"
	if config.ServerAddress != "" {
		httpAddr = config.ServerAddress
	}

	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: httpRouter,
	}

	service := &Service{
		Service: baseService,

		Config:     config,
		httpRouter: httpRouter,
		httpServer: httpServer,
	}
	return service, nil
}

func ProvideService(
	config *Config,
	baseService *base.Service,
) (*Service, func(), error) {
	service, err := NewService(config, baseService)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		service.Stop()
	}
	return service, cleanup, nil
}
