package application

import (
	"app/internal/ddd"
	"app/tenants/internal/domain"
	"context"
)

type CompanyHandlers[T ddd.AggregateEvent] struct {
	company domain.CompanyRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*CompanyHandlers[ddd.AggregateEvent])(nil)

func NewCompanyHandlers(company domain.CompanyRepository) *CompanyHandlers[ddd.AggregateEvent] {
	return &CompanyHandlers[ddd.AggregateEvent]{
		company: company,
	}
}

func (h CompanyHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.TenantCreatedEvent:
		return h.onTenantCreated(ctx, event)
	}
	return nil
}

func (h CompanyHandlers[T]) onTenantCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.TenantCreated)
	return h.company.AddTenant(ctx, event.AggregateID(), payload.Name)
}
