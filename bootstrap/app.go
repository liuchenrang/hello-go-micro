package bootstrap

import (
	"duoduo.com/micro/common/contracts"
)

var (
	App  = Application{boots:make([]contracts.Provider,0)}
)


type Application struct {
	boots        []contracts.Provider
}


func (app *Application) Register(boot contracts.Provider) {
	app.boots = append(app.boots, boot)
}

func (app *Application) Boot() {
	for _, value := range app.boots {
		value.Boot()
	}
}

func (app *Application) Shutdown() {
	for i := 0; i < len(app.boots); i++ {
		app.boots[i].Shutdown()
	}
}

func (app *Application) After() {
	for i := 0; i < len(app.boots); i++ {
		app.boots[i].After()
	}
}
