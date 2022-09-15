package main

import (
	"context"
	"flag"
	"log"
	"simple-storage/router"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	DefaultConfig = "config.yaml"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("[server] %s\n", err)
	}
}

func run() error {

	fileConfigPath := flag.String("config", DefaultConfig, " Configuration file")

	viper.SetConfigName("config")
	viper.AddConfigPath(*fileConfigPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[server] reading config file %s\n", err)
	}

	e := echo.New()
	e, err = router.Routes(e)
	if err != nil {
		return errors.Wrap(err, "creating router")
	}
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))

	timeoutCtx := 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), timeoutCtx)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
