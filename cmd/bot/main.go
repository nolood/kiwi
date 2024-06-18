package main

import (
	"kiwi/internal/config"
	"kiwi/internal/lib/logger"
)

func main() {

	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("Starting application...")

}
