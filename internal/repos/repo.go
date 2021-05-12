package repos

import (
	"go-wire-example/pkg/log"
)

type IRepository interface{}

type Repository struct {
	Logger log.ILogger
}

func NewRepository(
	logger log.ILogger,
) IRepository {
	repository := &Repository{
		Logger: logger,
	}
	return repository
}
