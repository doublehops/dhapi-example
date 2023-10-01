package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi-example/internal/config"
	"github.com/doublehops/dhapi-example/internal/db"
	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/middleware/database"
	"github.com/doublehops/dhapi-example/internal/middleware/logger"
	"github.com/doublehops/dhapi-example/internal/routes"
	"github.com/doublehops/dhapi-example/internal/runflags"
)

func main() {
	flags := runflags.GetFlags()
	// Setup config.
	cfg, err := config.New(flags.ConfigFile)
	if err != nil {
		log.Fatalf("error starting main. %s", err.Error())
	}

	// Setup logger.
	l := logga.New(&cfg.Logging)

	// Setup db connection.
	DB, err := db.New(l, cfg.DB)
	if err != nil {
		l.Error("error creating database connection. %s", err.Error())
		os.Exit(1)
	}

	// Setup and run Gin.
	r := gin.New()
	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("error setting trusted proxy. %s", err)
	}

	r.Use(gin.Recovery())
	r.Use(database.DatabaseMiddleware(DB))
	r.Use(logger.LoggingMiddleware(l))
	routes.GetRoutes(r)

	r.Run(":8080")
}
