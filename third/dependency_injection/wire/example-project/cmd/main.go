package main

import (
	"flag"
	"go-app/third/dependency_injection/wire/example-project/internal/config"
	"net/http"

	"go.uber.org/zap"
)

// Reference： https://gitee.com/huoyingwhw/kratos_study

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "../config/config.json", "config path, eg: -conf config.json")
}

func main() {
	flag.Parse()
	cfg := config.Load(configPath)
	app, cleanup, err := wireApp(cfg, cfg.Logger, cfg.Server)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Run()
}

type App struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	hs     *http.Server
}

func newApp(cfg *config.Config, logger *zap.SugaredLogger, hs *http.Server) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
		hs:     hs,
	}
}

func (a *App) Run() {
	if err := a.hs.ListenAndServe(); err != nil {
		panic(err)
	}
}
