package main

import (
	"app/internal/config"
	"app/internal/monolith"
	"app/internal/waiter"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	logger  zerolog.Logger
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

func (a *app) Logger() zerolog.Logger {
	return a.logger
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
	webServer := &http.Server{
		Addr:    a.cfg.Web.Address(),
		Handler: a.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Printf("web server started; listeningat http://localhost%s\n", a.cfg.Web.Port)
		defer fmt.Println("web server shutdown")
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}
