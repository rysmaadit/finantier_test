package service

import (
	"github.com/rysmaadit/finantier_test/auth_service/app"
	"github.com/rysmaadit/finantier_test/auth_service/external/jwt_client"
)

type Dependencies struct {
	AuthService AuthServiceInterface
}

func InstantiateDependencies(application *app.Application) Dependencies {
	jwtWrapper := jwt_client.New()
	authService := NewAuthService(application.Config, jwtWrapper)

	return Dependencies{
		AuthService: authService,
	}
}
