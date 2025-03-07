package queries

import (
	"app/tenants/internal/domain"
	"context"
)

type GetTenant struct {
	ID string
}

type GetCompanyHandler struct {
	company domain.CompanyRepository
}

func NewGetTenantHandler(company domain.CompanyRepository) GetCompanyHandler {
	return GetCompanyHandler{company: company}
}

func (h GetCompanyHandler) GetTenant(ctx context.Context, query GetTenant) (*domain.CompanyTenant, error) {
	return h.company.Find(ctx, query.ID)
}
