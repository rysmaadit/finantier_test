package service

import (
	"github.com/rysmaadit/finantier_test/encryption_service/app"
)

type Dependencies struct {
	EncryptionService EncryptionServiceInterface
}

func InstantiateDependencies(application *app.Application) Dependencies {
	authService := NewEncryptionService(application.Config)

	return Dependencies{
		EncryptionService: authService,
	}
}
