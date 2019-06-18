package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/webonise/csv_upload/app/configs"
	"github.com/webonise/csv_upload/app/containers"
	"github.com/webonise/csv_upload/pkg/bundle"
	"github.com/webonise/csv_upload/pkg/database"
	"github.com/webonise/csv_upload/pkg/envprovider"
	"github.com/webonise/csv_upload/pkg/logger"
	"github.com/webonise/csv_upload/pkg/router"
	"github.com/webonise/csv_upload/pkg/templates"
)

func main() {

	// Initialise environment variables
	v := &envprovider.EnvConfigProvider{}

	// Initialise server configurations
	cfg := &configs.ServerConfig{}
	cfg.InitialiseServerCfg(v)

	// Initialise logger
	log := &logger.RealLogger{}
	log.Initialise()

	d := &database.PGSQLDBConnectionInitialiser{
		DBConfig: &database.DBConfig{
			Username:         cfg.DBUsername,
			Password:         cfg.DBPassword,
			Host:             cfg.DBHost,
			DatabaseName:     cfg.DBName,
			ConnectionParams: cfg.DBConnParams,
			Port:             cfg.DBPORT,
		},
	}
	if err := d.InitialiseConnection(); err != nil {
		log.Panic(`Failed to create database connection`, err)
	}

	pc := &bundle.PackerConfig{
		Path: "../../migrations",
	}
	dbm := pc.Migration()
	if err := dbm.Migrate(d.Driver, d.DB); err != nil {
		panic(err)
	}

	pc.Path = "../../web"
	pc.Pack()

	// Initialise server application with essential dependencies
	srv := &containers.Server{
		Router: &router.Multiplexer{
			Mux: router.NewRouter(),
		},
		Cfg:       cfg,
		Log:       log,
		TplParser: &templates.TemplateParser{},
		DB:        d.DB,
	}

	srv.InitializeSrv()

	if err := http.ListenAndServe(cfg.Port, srv.Router.Mux); err != nil {
		panic(err)
	}
}
