package domain

import "context"

type CompanyTenant struct {
	ID   string
	Name string
}

type CompanyRepository interface {
	AddTenant(ctx context.Context, tenantID, name string) error
	Find(ctx context.Context, tenantID string) (*CompanyTenant, error)
}
