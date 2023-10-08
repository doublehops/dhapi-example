package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi-example/internal/config"
	"github.com/doublehops/dhapi-example/internal/db"
	"github.com/doublehops/dhapi-example/internal/handlers"
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

	app := &handlers.App{
		DB:  DB,
		Log: l,
	}

	// Setup and run Gin.
	r := gin.New()
	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return fmt.Errorf("error setting trusted proxy. %s", err)
	}

	r.Use(gin.Recovery())
	routes.GetRoutes(r, app)

	r.Run(":8080")

	return nil
}
