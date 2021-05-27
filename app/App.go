package app

type App struct {
	option *appOptions
}

func NewApp() *App {
	a := &App{}
	return a
}

func (app *App) Run() error {
	//start server
	err := app.option.server.Start()
	if err != nil {
		return err
	}
	//register

	return nil
}
