package main

import (
	"kiwi/internal/app"
	"kiwi/internal/config"
	"kiwi/internal/lib/logger"
	"log"
	"net/http"

	// _ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// startPprof()

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

func startPprof() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
