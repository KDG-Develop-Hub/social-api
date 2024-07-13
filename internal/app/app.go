package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kdg-develop-hub/api/config"
	v1 "github.com/kdg-develop-hub/api/internal/controller/http/v1"
	"github.com/kdg-develop-hub/api/internal/repository"
	"github.com/kdg-develop-hub/api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

type app struct {
	cfg *config.Config
}

type App interface {
	Run()
}

func NewApp(cfg *config.Config) App {
	return &app{cfg: cfg}
}

func (app *app) Run() {
	e := echo.New()

	// logger
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	// middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("method", v.Method).
				Msg("request")
			return nil
		},
	}))
	e.Use(middleware.Recover())

	// sql
	var driver string
	var source string

	switch app.cfg.Env {
	case config.Production:
		driver = "postgres"
		source = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			app.cfg.DB.Host,
			app.cfg.DB.Port,
			app.cfg.DB.User,
			app.cfg.DB.Password,
			app.cfg.DB.Name,
			app.cfg.DB.SslMode,
		)
	case config.Development:
		driver = "sqlite3"
		source = "db/development.sqlite3"
	}
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal().Err(err).Msg("DB close failed")
		}
	}()
	if err = db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("DB ping failed")
	}
	db = sqldblogger.OpenDriver(
		source,
		db.Driver(),
		zerologadapter.New(log),
	)
	dbx := sqlx.NewDb(db, driver)

	// router
	{
		ur := repository.NewUserRepository(dbx)

		us := service.NewUserService(ur)

		v1.NewRouter(e, &log, us)
	}

	// start server
	if err := e.Start(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}
