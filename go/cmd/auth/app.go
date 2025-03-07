package main

import (
	"app/internal/config"
	"app/internal/monolith"
	"app/internal/waiter"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

type app struct {
	cfg     config.AppConfig
	db      *pgx.Conn
	logger  zerolog.Logger
	modules []monolith.Module
	mux     *chi.Mux
	waiter  waiter.Waiter
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}

func (a *app) DB() *pgx.Conn {
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
		Addr: a.cfg.Web.Address(),
		// Handler: a.mux,
		// Use h2c so we can serve HTTP/2 without TLS.
		// TODO: Use TLS in production. See https://connectrpc.com/docs/go/deployment
		Handler: h2c.NewHandler(a.mux, &http2.Server{}),
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Printf("web server started; listening at http://localhost%s\n", a.cfg.Web.Port)
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
