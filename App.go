package toyms

type App struct {
	option *appOptions
}

func NewApp() *App {
	a := &App{}
	return a
}

func (a *App) Run() error {
	return nil
}
