package callbacks

import (
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/handlers/scenes"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/bot/static/texts"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"go.uber.org/zap"
)

type Callbacks struct {
	log      *zap.Logger
	services *services.Services
	bh       *th.BotHandler
	scenes   *scenes.Scenes
}

func New(log *zap.Logger, servs *services.Services, b *telego.Bot, bh *th.BotHandler, sc *scenes.Scenes) *Callbacks {

	return &Callbacks{
		log:      log,
		services: servs,
		bh:       bh,
		scenes:   sc,
	}
}

func (c *Callbacks) Register() {
	c.viewProfile()
	c.fillProfile()
}

func (c *Callbacks) viewProfile() {
	c.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

		chat := query.Message.GetChat()

		msg := telego.EditMessageTextParams{Text: "view profile", InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID()}

		_, err := bot.EditMessageText(&msg)
		if err != nil {
			c.log.Error("handlers.callbacks.ViewProfile", zap.Error(err))
		}

	}, th.CallbackDataEqual(callbacks_consts.VIEW_PROFILE))
}

func (c *Callbacks) fillProfile() {
	c.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		const op = "handlers.callbacks.FillProfile.Callback"

		chat := query.Message.GetChat()

		infoMsg := telego.EditMessageTextParams{Text: texts.ProfileFillInfo, InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID()}

		_, err := bot.EditMessageText(&infoMsg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

		c.scenes.Profile.StartFillProfileScene(chat.ChatID())

	}, th.CallbackDataEqual(callbacks_consts.FILL_PROFILE))

	// About handler
	// bh.Handle(func(bot *telego.Bot, update telego.Update) {
	// 	const op = "handlers.callbacks.FillProfile.Message.About"

	// 	about := update.Message.Text
	// 	var msg *telego.SendMessageParams

	// 	c.log.Info(about)

	// 	keyboard := tu.InlineKeyboard(tu.InlineKeyboardRow(
	// 		tu.InlineKeyboardButton(texts.GenderMale).WithCallbackData(GENDER_MALE),
	// 		tu.InlineKeyboardButton(texts.GenderFemale).WithCallbackData(GENDER_FEMALE),
	// 	))

	// 	msg = tu.Message(
	// 		tu.ID(update.Message.Chat.ID),
	// 		"Complete",
	// 	).WithReplyMarkup(keyboard)

	// 	c.services.Session.Set(update.Message.From.ID, model.Session_None)
	// 	c.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{About: &about})

	// 	_, err := bot.SendMessage(msg)
	// 	if err != nil {
	// 		c.log.Error(op, zap.Error(err))
	// 	}

	// }, th.And(th.AnyMessageWithText(), predicates.ThMessageSessionEqual(*c.services, model.Session_FillProfileAbout)))

}
