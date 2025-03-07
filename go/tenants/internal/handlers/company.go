package handlers

import (
	"app/internal/ddd"
	"app/tenants/internal/domain"
)

func RegisterTenantsHandlers(tenantsHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.TenantCreatedEvent, tenantsHandlers)
}
