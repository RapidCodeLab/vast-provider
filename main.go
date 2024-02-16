package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RapidCodeLab/vast-provider/provider"
	"github.com/caarlos0/env"
	"golang.org/x/exp/slog"
)

const timeFormat = "2006-01-02 15:04:05"

type (
	Config struct {
		ServerListenNetwork string `env:"SERVER_LISTEN_NETWORK,required"`
		ServerListenAddr   string `env:"SERVER_LISTEN_ADDR,required"`
		BaseURL            string `env:"BASE_URL,required"`
	}
)

func main() {
	slog.Info(
		"VAST Provider Started",
		"datetim", time.Now().Format(timeFormat),
	)

	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		slog.Error(
			"Parsing Config Error",
			"datetime", time.Now().Format(timeFormat),
			"error", err.Error(),
		)
		os.Exit(1)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		defer signal.Stop(exit)
		<-exit
		cancel()
	}()

	p := provider.New()

	err = p.Start(
		ctx,
		config.ServerListenNetwork,
		config.ServerListenAddr,
		config.BaseURL,
	)
	if err != nil {
		slog.Error(
			"VAST ProviderStopped With Error",
			"datetime", time.Now().Format(timeFormat),
			"error", err.Error(),
		)
	}

	slog.Info(
		"VAST Provider Successfully Stopped",
		"datetime", time.Now().Format(timeFormat),
	)
}
