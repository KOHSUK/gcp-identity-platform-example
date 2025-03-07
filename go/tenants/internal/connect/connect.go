package connect

import (
	v1 "app/internal/gen/proto/tenants/tenantspb/v1"
	"app/internal/gen/proto/tenants/tenantspb/v1/tenantsv1connect"
	"app/tenants/internal/application"
	"app/tenants/internal/application/commands"
	"app/tenants/internal/application/queries"
	"context"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TenantsServer struct {
	app application.App
}

func (s *TenantsServer) CreateTenant(ctx context.Context, request *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	tenantID := uuid.New().String()

	err := s.app.CreateTenant(ctx, commands.CreateTenant{
		ID:   tenantID,
		Name: request.Msg.Name,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.CreateTenantResponse{
		Id: tenantID,
	}), nil
}

func (s *TenantsServer) GetTenant(ctx context.Context, request *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error) {
	tenant, err := s.app.GetTenant(ctx, queries.GetTenant{ID: request.Msg.Id})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.GetTenantResponse{
		Tenant: &v1.Tenant{
			Id:   tenant.ID,
			Name: tenant.Name,
		},
	}), nil
}

func RegisterConnect(ctx context.Context, mux *chi.Mux, app application.App) error {
	server := &TenantsServer{
		app: app,
	}
	path, handler := tenantsv1connect.NewTenantsServiceHandler(server)
	mux.Mount(path, handler)
	return nil
}
