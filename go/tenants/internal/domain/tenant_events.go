package domain

const (
	TenantCreatedEvent = "tenants.TenantCreated"
)

type TenantCreated struct {
	Name string
}

// Key implements registry.Registerable
func (TenantCreated) Key() string { return TenantCreatedEvent }
