package callbacks

import (
	"kiwi/.gen/kiwi/public/model"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/bot/static/texts"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

const (
	VIEW_PROFILE  = "view_profile"
	FILL_PROFILE  = "fill_profile"
	GENDER_MALE   = "gender_male"
	GENDER_FEMALE = "gender_female"
)

type Callbacks interface {
	ViewProfile(bh *th.BotHandler)
	FillProfile(bh *th.BotHandler)
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

func (c *callbacks) FillProfile(bh *th.BotHandler) {

	// Start profile handler

	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		const op = "handlers.callbacks.FillProfile.Callback"

		chat := query.Message.GetChat()
		msg := telego.EditMessageTextParams{Text: texts.AgeQuestion, InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID()}

		c.services.Session.Set(query.From.ID, model.Session_FillProfileAge)

		_, err := bot.EditMessageText(&msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.CallbackDataEqual(FILL_PROFILE))

	// Age handler

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		const op = "handlers.callbacks.FillProfile.Message.Age"

		session, err := c.services.Session.Get(update.Message.From.ID)
		if err != nil {
			c.log.Error(op, zap.Error(err))
			return
		}

		if session != model.Session_FillProfileAge {
			return
		}

		textAge := update.Message.Text
		var msg *telego.SendMessageParams

		age, _ := strconv.Atoi(textAge)

		ok := true

		keyboard := tu.InlineKeyboard(tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.GenderMale).WithCallbackData(GENDER_MALE),
			tu.InlineKeyboardButton(texts.GenderFemale).WithCallbackData(GENDER_FEMALE),
		))

		msg = tu.Message(
			tu.ID(update.Message.Chat.ID),
			texts.GenderQuestion,
		).WithReplyMarkup(keyboard)

		if age < 16 {
			ok = false
			msg = tu.Message(
				tu.ID(update.Message.Chat.ID),
				texts.AgeLower,
			)
		}

		if age > 150 {
			ok = false
			msg = tu.Message(
				tu.ID(update.Message.Chat.ID),
				texts.AgeUpper,
			)
		}

		_, err = strconv.Atoi(textAge)
		if err != nil {
			ok = false
			msg = tu.Message(
				tu.ID(update.Message.Chat.ID),
				texts.AgeIncorrect,
			)
		}

		if ok {
			c.services.Session.Set(update.Message.From.ID, model.Session_FillProfileGender)
		}

		_, err = bot.SendMessage(msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.AnyMessageWithText())

}
