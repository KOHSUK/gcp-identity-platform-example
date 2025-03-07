package monolith

import (
	"app/internal/config"
	"app/internal/waiter"
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *pgx.Conn
	Logger() zerolog.Logger
	Mux() *chi.Mux
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
