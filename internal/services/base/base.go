package base

import (
	"go-wire-example/internal/repos"
	"go-wire-example/pkg/log"
)

type Service struct {
	Logger     log.ILogger
	Repository repos.IRepository
}

func (s *Service) Start() {

}

func (s *Service) Stop() {

}

func NewService(
	logger log.ILogger,
	repository repos.IRepository,
) (*Service, error) {
	baseService := &Service{
		Logger:     logger,
		Repository: repository,
	}
	return baseService, nil
}

func ProvideService(
	logger log.ILogger,
	repository repos.IRepository,
) (*Service, func(), error) {
	baseService, err := NewService(logger, repository)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		baseService.Stop()
	}
	return baseService, cleanup, nil
}
