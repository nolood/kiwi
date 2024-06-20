package commands

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

type Commands interface {
}

func Register(log *zap.Logger, bh *th.BotHandler) {
	startCommand(bh)
}

func startCommand(bh *th.BotHandler) {
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Hello, "+update.Message.From.FirstName,
		))
	}, th.CommandEqual("start"))
}
