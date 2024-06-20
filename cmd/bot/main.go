package main

import (
	"kiwi/internal/app"
	"kiwi/internal/config"
	"kiwi/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	application := app.New(log, cfg)

	go application.Bot.MustRun()

	log.Info("App started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Bot.Stop()

	log.Info("App stopped")
}
