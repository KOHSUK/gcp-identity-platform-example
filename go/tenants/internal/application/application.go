package application

import (
	"app/tenants/internal/application/commands"
	"app/tenants/internal/application/queries"
	"app/tenants/internal/domain"
	"context"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateTenant(ctx context.Context, cmd commands.CreateTenant) error
	}
	Queries interface {
		GetTenant(ctx context.Context, query queries.GetTenant) (*domain.CompanyTenant, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateTenantHandler
	}
	appQueries struct {
		queries.GetCompanyHandler
	}
)

var _ App = (*Application)(nil)

func New(tenants domain.TenantRepository, companies domain.CompanyRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateTenantHandler: commands.NewCreateTenantHandler(tenants),
		},
		appQueries: appQueries{
			GetCompanyHandler: queries.NewGetTenantHandler(companies),
		},
	}
}
