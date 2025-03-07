package domain

import "context"

type TenantRepository interface {
	Load(ctx context.Context, tenantID string) (*Tenant, error)
	Save(ctx context.Context, tenant *Tenant) error
}
