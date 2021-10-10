package main

import (
	"github.com/Eri-Vadi/go_kitchen/application"
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/config"
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/profiler"
	"os"
	"os/signal"
	"syscall"
)

func init(){
	config.LoadConfig()
}

func main(){
	profiler.StartProfiler()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	app := application.CreateApp()
	go app.Start()

	<-sigChan
}
