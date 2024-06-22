package bot

import (
	"kiwi/internal/app/bot/handlers"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/config"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type BotApp struct {
	log      *zap.Logger
	cfg      config.Telegram
	Bot      *telego.Bot
	services *services.Services
}

func New(log *zap.Logger, cfg config.Telegram, servs *services.Services) *BotApp {
	const op = "bot.New"

	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Panic(op, zap.Error(err))
	}

	return &BotApp{
		log:      log,
		cfg:      cfg,
		services: servs,
		Bot:      bot,
	}
}

func (b *BotApp) MustRun() {
	const op = "bot.MustRun"

	updates, err := b.Bot.UpdatesViaLongPolling(nil)
	if err != nil {
		b.log.Error(op, zap.Error(err))
	}

	handlers.Register(b.log, updates, b.Bot, b.services)

	defer b.Bot.StopLongPolling()
}

func (b *BotApp) Stop() {
	b.Bot.StopLongPolling()
}
