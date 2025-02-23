package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"yadro.com/course/internal/apiserver"
	"yadro.com/course/internal/logger"
)

var log = logger.GetInstance()

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
	done := make(chan bool)

	srv := apiserver.CreateServer()
	// srv.GracefulShutDown()
	go shutDownListener(srv, done)
	if err := srv.Run(); err != nil {
		log.Error("Server run error", err)
		panic(err) //TODO: handle error
	}

	<-done // Ждем завершения
}
