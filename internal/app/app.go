package app

import (
	botapp "kiwi/internal/app/bot"
	"kiwi/internal/config"
	"kiwi/internal/storage/postgres"

	"go.uber.org/zap"
)

type App struct {
	Bot *botapp.BotApp
}

func New(log *zap.Logger, cfg *config.Config) *App {

	_ = postgres.New(cfg.Storage)

	bot := botapp.New(cfg.Telegram)

	return &App{
		Bot: bot,
	}
}
