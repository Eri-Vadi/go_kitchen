package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Eri-Vadi/go_kitchen/application"
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	app := application.CreateApp()
	go app.Start()

	<-sigChan

	app.Shutdown(ctx)
	cancel()
}
