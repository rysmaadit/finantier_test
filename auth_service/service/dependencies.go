package service

import (
	"github.com/rysmaadit/finantier_test/auth_service/app"
)

type Dependencies struct {
	AuthService AuthServiceInterface
}

func InstantiateDependencies(application *app.Application) Dependencies {
	authService := NewAuthService(application.Config)

	return Dependencies{
		AuthService: authService,
	}
}
