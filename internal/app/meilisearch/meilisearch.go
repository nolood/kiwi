package meilisearch

import (
	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
	"kiwi/internal/config"
)

type App struct {
	log    *zap.Logger
	cfg    config.Meilisearch
	Client *meilisearch.Client
}

func New(log *zap.Logger, cfg config.Meilisearch) *App {

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   cfg.Address,
		APIKey: cfg.Key,
	})

	return &App{
		log:    log,
		cfg:    cfg,
		Client: client,
	}
}
