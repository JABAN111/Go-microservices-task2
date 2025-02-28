package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	"yadro.com/course/internal/apiserver"
	"yadro.com/course/internal/config"
	"yadro.com/course/internal/logger"
)

var log = logger.GetInstance()

const defaultConfigPath = "config.yaml"

var configPath string

func getConfig() *config.Config {
	cfg := config.NewConfig()

	if err := cleanenv.ReadConfig(configPath, cfg); err == nil {
		if cfg.BindPort != "" {
			log.Info(fmt.Sprintf("Using configuration: %s", cfg))
			return cfg
		}
	}

	log.Warn("Failed to load configuration from file or environment variables")

	return config.DefaultConfig()
}

func shutDownListener(s apiserver.ServerManipulator, done chan<- bool) {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()
	<-ctx.Done()

	if err := s.GracefulShutDown(); err != nil {
		log.Error("Failed while shutting down the server %s", slog.Any("err", err))
	}
	log.Info("Write for closing channel...")
	done <- true
}

func main() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to config file")
	flag.Parse()

	cfg := getConfig()

	done := make(chan bool)

	port, err := strconv.Atoi(cfg.BindPort)
	if err != nil {
		log.Error("Cannot parse port to integer, using default port value")
		port = 8080
	}
	srv := apiserver.CreateServer(port)

	go shutDownListener(srv, done)
	if err := srv.Run(); err != nil {
		log.Error(fmt.Sprintf("Server run error: %s", err.Error()))
		panic(err)
	}

	<-done
}
