package bot

import (
	"kiwi/internal/app/bot/handlers"
	"kiwi/internal/config"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type BotApp struct {
	log     *zap.Logger
	cfg     config.Telegram
	Updates <-chan telego.Update
	Bot     *telego.Bot
}

func New(log *zap.Logger, cfg config.Telegram) *BotApp {
	const op = "bot.New"

	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Panic(op, zap.Error(err))
	}

	return &BotApp{
		log: log,
		cfg: cfg,
		Bot: bot,
	}
}

func (b *BotApp) MustRun() {
	const op = "bot.MustRun"

	updates, err := b.Bot.UpdatesViaLongPolling(nil)
	if err != nil {
		b.log.Error(op, zap.Error(err))
	}

	handlers.Register(b.log, updates, b.Bot)

	defer b.Bot.StopLongPolling()
}

func (b *BotApp) Stop() {
	b.Bot.StopLongPolling()
}
