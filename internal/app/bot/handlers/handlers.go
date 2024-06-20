package handlers

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"

	"go.uber.org/zap"
)

func Register(log *zap.Logger, updates <-chan telego.Update, b *telego.Bot) {
	bh, _ := th.NewBotHandler(b, updates)

	startCommand(bh)

	bh.Start()
}

func startCommand(bh *th.BotHandler) {
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Hello, "+update.Message.From.FirstName,
		))
	}, th.CommandEqual("start"))
}
