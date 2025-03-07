package domain

import (
	"app/internal/ddd"
	"app/internal/es"
	"errors"
	"fmt"
)

const TenantAggregate = "tenants.Tenant"

var (
	ErrTenantNameIsBlank = errors.New("the tenant name cannot be blank")
)

type Tenant struct {
	es.Aggregate
	Name string
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Tenant)(nil)

func NewTenant(id string) *Tenant {
	return &Tenant{
		Aggregate: es.NewAggregate(id, TenantAggregate),
	}
}

func CreateTenant(id, name string) (*Tenant, error) {
	if name == "" {
		return nil, ErrTenantNameIsBlank
	}

	tenant := NewTenant(id)

	tenant.AddEvent(TenantCreatedEvent, &TenantCreated{
		Name: name,
	})

	return tenant, nil
}

// Key implements registry.Registerable
func (Tenant) Key() string { return TenantAggregate }

// ApplyEvent implements es.EventApplier
func (t *Tenant) ApplyEvent(event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *TenantCreated:
		t.Name = payload.Name
	default:
		return fmt.Errorf("%T received the event %s with an unexpected payload %T", t, event.EventName(), payload)
	}

	return nil
}

// Snapshot implements es.Snapshotter
func (t *Tenant) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *TenantV1:
		t.Name = ss.Name
	default:
		return fmt.Errorf("%T received the snapshot with an unexpected type %T", t, snapshot)
	}

	return nil
}

// ToSnapshot implements es.Snapshotter
func (t Tenant) ToSnapshot() es.Snapshot {
	return TenantV1{
		Name: t.Name,
	}
}
