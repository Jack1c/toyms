package app

import "github.com/Jack1c/toyms/log"

type App struct {
	option *appOption
}

var defaultOption = &appOption{
	logger: log.DefaultLogger,
}

func NewApp(opt ...AppOptions) *App {
	option := defaultOption
	for _, o := range opt {
		o.apply(option)
	}

	a := &App{
		option: option,
	}
	return a
}

func (app *App) Run() error {
	//start server
	app.option.logger.Log(log.LevelInfo, "msg", "server start ")
	err := app.option.server.Start()
	if err != nil {
		return err
	}
	//register
	app.option.logger.Log(log.LevelInfo, "msg", "server register ")

	return nil
}
