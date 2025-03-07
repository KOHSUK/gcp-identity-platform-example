package logging

import (
	"app/tenants/internal/application"
	"app/tenants/internal/application/commands"
	"app/tenants/internal/application/queries"
	"app/tenants/internal/domain"
	"context"

	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) CreateTenant(ctx context.Context, cmd commands.CreateTenant) (err error) {
	a.logger.Info().Msg("--> Tenants.CreateTenant")
	defer func() { a.logger.Info().Err(err).Msg("<-- Tenants.CreateTenant") }()
	return a.App.CreateTenant(ctx, cmd)
}

func (a Application) GetTenant(ctx context.Context, query queries.GetTenant) (tenant *domain.CompanyTenant, err error) {
	a.logger.Info().Msg("--> Tenants.GetTenant")
	defer func() { a.logger.Info().Err(err).Msg("<-- Tenants.GetTenant") }()
	return a.App.GetTenant(ctx, query)
}
