package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/config"
	"github.com/doublehops/dhapi-example/internal/db"
	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/routes"
	"github.com/doublehops/dhapi-example/internal/runflags"
)

func main() {
	if err := run(); err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func run() error {
	flags := runflags.GetFlags()

	// Setup config.
	cfg, err := config.New(flags.ConfigFile)
	if err != nil {
		return fmt.Errorf("error starting main. %s", err.Error())
	}

	// Setup logger.
	l, err := logga.New(&cfg.Logging)
	if err != nil {
		return fmt.Errorf("error configuring logger. %s", err.Error())
	}

	// Setup db connection.
	DB, err := db.New(l, cfg.DB)
	if err != nil {
		return fmt.Errorf("error creating database connection. %s", err.Error())
	}

	App := &app.App{
		DB:  DB,
		Log: l,
	}

	router := httprouter.New()
	routes := routes.GetV1Routes(App)

	for _, r := range routes.Routes() {
		fmt.Printf(">>> %s %s\n", r.Method(), r.Path())
		router.Handle(r.Method(), r.Path(), r.Handler())
	}

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return fmt.Errorf("unable to start server. %s", err.Error())
	}

	return nil
}
