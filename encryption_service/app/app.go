package app

import (
	"github.com/rysmaadit/finantier_test/encryption_service/config"
)

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
