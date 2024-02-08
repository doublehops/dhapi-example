package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/doublehops/dhapi-example/internal/config"
	"github.com/doublehops/dhapi-example/internal/db"
	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/scaffold"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	migrate "github.com/doublehops/go-migration"
)

/*
  This file just serves as an example of how you would add this library to your project.
*/

func main() {
	if err := run(); err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func run() error {
	var args migrate.Action

	flag.StringVar(&args.Action, "model", "", "the database table/model to scaffold")

	configFile := flag.String("config", "config.json", "Scaffold file to use")
	flag.Parse()

	// Setup config.
	cfg, err := config.New(*configFile)
	if err != nil {
		return fmt.Errorf("error starting. %s", err.Error())
	}

	scf, err := GetScaffoldConfig()
	if err != nil {
		return fmt.Errorf("error getting scaffolding config. %s", err.Error())
	}

	scf.Run()

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

	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("there was an error with os.Getwd(). %s", err.Error())
	}

	args.Path = dir + "/migrations"
	args.DB = DB

	err = args.Migrate()
	if err != nil {
		return fmt.Errorf("there was an error initialising migration. %s", err.Error())
	}

	return nil
}

func GetScaffoldConfig() (*scaffold.Scaffold, error) {
	sc := &scaffold.Scaffold{}
	pwd, err := os.Getwd()
	if err != nil {
		return sc, fmt.Errorf("there was an error with os.Getwd(). %s", err.Error())
	}

	relPath := pwd + "/config.json"

	f, err := os.ReadFile(relPath)
	if err != nil {
		log.Printf("unable to read config file - %s. %s", relPath, err.Error())

		return nil, fmt.Errorf("unable to read config file `%s`. %w", relPath, err)
	}

	if err = json.Unmarshal(f, sc); err != nil {
		return sc, err
	}

	return sc, nil
}
