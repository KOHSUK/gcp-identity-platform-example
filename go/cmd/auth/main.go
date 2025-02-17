package main

import (
	"app/internal/config"
	"app/internal/logger"
	"app/internal/monolith"
	"app/internal/waiter"
	"app/internal/web"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	// init infrastructure...
	// init db
	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)
	m.logger = initLogger(cfg)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	m.modules = []monolith.Module{}

	if err = m.startupModules(); err != nil {
		return err
	}

	fmt.Println("started application")
	defer fmt.Println("stopped application")

	m.waiter.Add(m.waitForWeb)

	return m.waiter.Wait()
}

func initLogger(cfg config.AppConfig) zerolog.Logger {
	return logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
