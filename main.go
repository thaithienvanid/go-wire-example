package main

import (
	_ "go.uber.org/automaxprocs"

	"go-wire-example/internal/services/rest_api"
	"go-wire-example/pkg/log"
	"go-wire-example/pkg/signal"
)

type Application struct {
	Config         *Config
	Logger         log.ILogger
	RestAPIService *rest_api.Service
}

func (a *Application) Start() {
	a.Logger.Info("app is starting")

	a.RestAPIService.Start()
	a.Logger.Info("app is running")

	s := signal.WaitOSSignal()
	a.Logger.Infof("received signal: %s", s.String())
}

func (a *Application) Stop() {
	a.Logger.Info("app is stopping")
	//
	a.Logger.Info("app is stopped!")
}

func NewApplication(
	config *Config,
	logger log.ILogger,
	restAPIService *rest_api.Service,
) (*Application, error) {
	app := &Application{
		Config:         config,
		Logger:         logger,
		RestAPIService: restAPIService,
	}
	return app, nil
}

func ProvideApplication(
	config *Config,
	logger log.ILogger,
	restAPIService *rest_api.Service,
) (*Application, func(), error) {
	app, err := NewApplication(
		config,
		logger,
		restAPIService,
	)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		app.Logger.Info("app is cleaning")
		app.Stop()
		app.Logger.Info("app is cleaned!")
	}
	return app, cleanup, nil
}

func main() {
	app, cleanup, err := InitializeApplication()
	if err != nil {
		app.Logger.Fatal(err)
		panic(err)
	}
	defer cleanup()
	app.Start()
}
