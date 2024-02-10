package main

import (
	"context"
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
	var modelName string

	flag.StringVar(&modelName, "model", "", "the database table/model to scaffold")

	configFile := flag.String("config", "config.json", "Scaffold file to use")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("there was an error with os.Getwd(). %s", err.Error())
	}

	// Setup config.
	cfg, err := config.New(*configFile)
	if err != nil {
		return fmt.Errorf("error starting. %s", err.Error())
	}

	// Setup logger.
	l, err := logga.New(&cfg.Logging)
	if err != nil {
		return fmt.Errorf("error configuring logger. %s", err.Error())
	}

	l.Info(context.TODO(), "flags", logga.KVPs{"model": modelName})

	// Setup db connection.
	DB, err := db.New(l, cfg.DB)
	if err != nil {
		return fmt.Errorf("error creating database connection. %s", err.Error())
	}

	//dir, err := os.Getwd()
	//if err != nil {
	//	return fmt.Errorf("there was an error with os.Getwd(). %s", err.Error())
	//}

	scf, err := GetScaffoldConfig(pwd)
	if err != nil {
		return fmt.Errorf("error getting scaffold config. %s", err.Error())
	}

	s := scaffold.New(pwd, scf, modelName, DB, l)
	err = s.Run()
	if err != nil {
		return err
	}

	return nil
}

func GetScaffoldConfig(pwd string) (scaffold.Config, error) {
	cp := scaffold.Config{}

	relPath := pwd + "/cmd/scaffold/config.json"

	f, err := os.ReadFile(relPath)
	if err != nil {
		log.Printf("unable to read config file - %s. %s", relPath, err.Error())

		return cp, fmt.Errorf("unable to read config file `%s`. %w", relPath, err)
	}

	if err = json.Unmarshal(f, &cp); err != nil {
		return cp, err
	}

	return cp, nil
}
