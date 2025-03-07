package commands

import (
	"app/tenants/internal/domain"
	"context"
)

type (
	CreateTenant struct {
		ID   string
		Name string
	}

	CreateTenantHandler struct {
		tenants domain.TenantRepository
	}
)

func NewCreateTenantHandler(tenants domain.TenantRepository) CreateTenantHandler {
	return CreateTenantHandler{tenants: tenants}
}

func (h CreateTenantHandler) CreateTenant(ctx context.Context, cmd CreateTenant) error {
	tenant, err := domain.CreateTenant(cmd.ID, cmd.Name)
	if err != nil {
		return err
	}

	return h.tenants.Save(ctx, tenant)
}
