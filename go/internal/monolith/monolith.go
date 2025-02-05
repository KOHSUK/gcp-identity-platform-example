package monolith

import (
	"app/internal/config"
	"app/internal/waiter"
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	Mux() *chi.Mux
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
