package tenants

import (
	"app/internal/ddd"
	"app/internal/es"
	"app/internal/monolith"
	pg "app/internal/postgres"
	"app/internal/registry"
	"app/internal/registry/serdes"
	"app/tenants/internal/application"
	"app/tenants/internal/connect"
	"app/tenants/internal/domain"
	"app/tenants/internal/handlers"
	"app/tenants/internal/logging"
	"app/tenants/internal/postgres"
	"context"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("tenants.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("tenants.snapshots", mono.DB(), reg),
	)
	tenants := es.NewAggregateRepository[*domain.Tenant](domain.TenantAggregate, reg, aggregateStore)
	company := postgres.NewCompanyRepository("tenants.tenants", mono.DB())

	app := logging.LogApplicationAccess(
		application.New(tenants, company),
		mono.Logger(),
	)

	companyHandlers := logging.LogEventHandlerAccess(
		application.NewCompanyHandlers(company),
		"Company", mono.Logger(),
	)

	if err := connect.RegisterConnect(ctx, mono.Mux(), app); err != nil {
		return err
	}

	handlers.RegisterTenantsHandlers(companyHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Tenant
	if err = serde.Register(domain.Tenant{}, func(v any) error {
		tenant := v.(*domain.Tenant)
		tenant.Aggregate = es.NewAggregate("", domain.TenantAggregate)
		return nil
	}); err != nil {
		return
	}
	// Tenant events
	if err = serde.Register(domain.TenantCreated{}); err != nil {
		return
	}
	// Tenant snashots
	if err = serde.RegisterKey(domain.TenantV1{}.SnapshotName(), domain.TenantV1{}); err != nil {
		return
	}

	return
}
