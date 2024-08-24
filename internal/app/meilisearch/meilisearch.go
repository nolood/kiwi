package meilisearch

import (
	"kiwi/internal/app/meilisearch/services"
	"kiwi/internal/config"

	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
)

type App struct {
	log      *zap.Logger
	Client   *meilisearch.Client
	cfg      config.Meilisearch
	Services *services.Services
}

func New(log *zap.Logger, cfg config.Meilisearch) *App {

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   cfg.Address,
		APIKey: cfg.Key,
	})

	servs := services.New(log, client)

	return &App{
		log:      log,
		cfg:      cfg,
		Client:   client,
		Services: servs,
	}
}
