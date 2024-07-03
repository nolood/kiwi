package callbacks

import (
	"kiwi/internal/app/bot/services"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"go.uber.org/zap"
)

const (
	VIEW_PROFILE = "view_profile"
)

type Callbacks interface {
	ViewProfile(bh *th.BotHandler)
}

type callbacks struct {
	log      *zap.Logger
	services *services.Services
}

func New(log *zap.Logger, servs *services.Services) Callbacks {
	return &callbacks{
		log:      log,
		services: servs,
	}
}

func (c *callbacks) ViewProfile(bh *th.BotHandler) {
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

		chat := query.Message.GetChat()

		msg := telego.EditMessageTextParams{Text: "view profile", InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID()}

		_, err := bot.EditMessageText(&msg)
		if err != nil {
			c.log.Error("handlers.callbacks.ViewProfile", zap.Error(err))
		}

	}, th.CallbackDataEqual(VIEW_PROFILE))
}
