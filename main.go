package main

import (
	"os"

	"creativecapsule/handlers"

	database "creativecapsule/db"

	"github.com/ardanlabs/conf"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}

}

func run() error {
	// Configuration

	cfg := struct {
		conf.Version
		Environment           string `conf:"default:local"`
		ContainerInternalPort string `conf:"default:9876"`
		AppHost               string `conf:"default:localhost"`
		AppProtocol           string `conf:"default:http"`
		PDB                   struct {
			Uri      string `conf:""`
			User     string `conf:"default:creativecapsule"`
			Password string `conf:"default:creativecapsule,noprint"`
			Host     string `conf:"default:localhost"`
			Port     string `conf:"default:5432"`
			Name     string `conf:"default:creativecapsuledb"`
			SslMode  string `conf:"default:disable"`
		}
		LogLevel string `conf:"default:trace"`
	}{}

	cfg.Version.Desc = "all rights reserved by adlevo tech"

	// ======================================================================
	// configuration initialization
	if err := conf.Parse(os.Args[1:], "CORE", &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, _ := conf.Usage("CORE", cfg)
			log.Info().Msg(usage)
			os.Exit(3)
		case conf.ErrVersionWanted:
			version, _ := conf.VersionString("CORE", cfg)
			log.Info().Msg(version)
			os.Exit(3)
		}
		log.Error().Err(err).Msg("")
		os.Exit(3)
	}

	// Log setup
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999Z07:00"
	zerolog.LevelFieldName = "level_name"
	logger := zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()
	e := echo.New()
	e.Use(middleware.AddTrailingSlash())

	var pdb *sqlx.DB
	var err error
	if cfg.PDB.Uri != "" {
		pdb, err = database.GetSQLXPGInstanceFromDsn(cfg.PDB.Uri)
		if err != nil {
			return errors.Wrap(err, "starting api server")
		}
	} else {
		pdb, err = database.GetSQLXPGInstance(cfg.PDB.User, cfg.PDB.Password, cfg.PDB.Host, cfg.PDB.Port, cfg.PDB.Name, cfg.PDB.SslMode)
		if err != nil {
			return errors.Wrap(err, "starting api server")
		}
	}

	defer func() {
		log.Printf("main: Database stopping : %s", cfg.PDB.Host)
		_ = pdb.Close()
	}()

	e.Use(middleware.Recover())
	handlers.API(e, logger, pdb)
	e.HideBanner = true
	e.HidePort = true

	return e.Start(":1323")
}
