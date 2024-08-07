package app

import (
	botapp "kiwi/internal/app/bot"
	"kiwi/internal/app/bot/repositories"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/meilisearch"
	"kiwi/internal/config"
	"kiwi/internal/storage/postgres"

	"go.uber.org/zap"
)

type App struct {
	Bot     *botapp.App
	MClient *meilisearch.App
}

func New(log *zap.Logger, cfg *config.Config) *App {

	mClient := meilisearch.New(log, cfg.Meilisearch)

	storage := postgres.New(cfg.Storage)

	repos := repositories.New(log, storage)

	servs := services.New(log, repos)

	bot := botapp.New(log, cfg.Telegram, servs)

	return &App{
		Bot:     bot,
		MClient: mClient,
	}
}
