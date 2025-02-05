package main

import (
	"app/internal/config"
	"app/internal/monolith"
	"app/internal/waiter"
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	modules []monolith.Module
	mux     *chi.Mux
	waiter  waiter.Waiter
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}

func (a *app) DB() *sql.DB {
	return a.db
}

func (a *app) Mux() *chi.Mux {
	return a.mux
}

func (a *app) Waiter() waiter.Waiter {
	return a.waiter
}

func (a *app) startupModules() error {
	for _, module := range a.modules {
		if err := module.Startup(a.Waiter().Context(), a); err != nil {
			return err
		}
	}

	return nil
}

func (a *app) waitForWeb(ctx context.Context) error {
	webServer := &http.Server{}
}
