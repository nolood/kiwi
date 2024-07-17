package main

import (
	"kiwi/internal/app/meilisearch"
	"kiwi/internal/config"
	"kiwi/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	meili := meilisearch.New(log, cfg.Meilisearch)

}
