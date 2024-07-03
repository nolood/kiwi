package commands

import (
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/bot/static/texts"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

type Commands interface {
	Start(bh *th.BotHandler)
}

type commands struct {
	log      *zap.Logger
	services *services.Services
}

func New(log *zap.Logger, servs *services.Services) Commands {
	return &commands{
		log:      log,
		services: servs,
	}
}

func (c *commands) Start(bh *th.BotHandler) {
	bh.Handle(func(bot *telego.Bot, update telego.Update) {

		user, err := c.services.User.GetOrCreate(update.Message.From)
		if err != nil {
			c.log.Error("handlers.commands.Start", zap.Error(err))
			return
		}

		c.log.Info("handlers.commands.Start", zap.Any("user", user))

		// TODO: Если есть анкета и отключена - | Начать поиск | Посмотреть анкету |
		// TODO: Если нет анкеты | Заполнить анкету |
		// TODO: Если есть анкета и включена - | Продолжить поиск | Посмотреть анкету |

		keyboard := tu.InlineKeyboard(tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Start").WithCallbackData("start"),
		))

		msg := tu.Message(
			tu.ID(update.Message.Chat.ID),
			texts.GreetingText(update.Message.From.FirstName),
		).WithReplyMarkup(keyboard)

		_, _ = bot.SendMessage(msg)

	}, th.CommandEqual("start"))
}
