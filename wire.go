//+build wireinject

package main

import (
	_ "github.com/google/subcommands"
	"github.com/google/wire"

	"go-wire-example/internal/repos"
	"go-wire-example/internal/services/base"
	"go-wire-example/internal/services/rest_api"
	"go-wire-example/pkg/log"
)

var CoreSet = wire.NewSet(
	ProvideConfig,
	log.ProvideLogger,
	repos.NewRepository,
	base.NewService,
)

var RestAPISet = wire.NewSet(
	ProvideRestAPIConfig,
	rest_api.NewService,
)

func InitializeApplication() (*Application, func(), error) {
	wire.Build(
		CoreSet,
		RestAPISet,
		ProvideApplication,
	)
	return &Application{}, nil, nil
}
