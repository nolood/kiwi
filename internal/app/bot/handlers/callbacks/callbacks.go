package callbacks

import (
	"kiwi/.gen/kiwi/public/model"
	userdto "kiwi/internal/app/bot/dto/user"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/bot/static/texts"
	"kiwi/internal/app/bot/utils/predicates"
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
	DEFAULT_PHOTO = "default_photo"
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

		infoMsg := telego.EditMessageTextParams{Text: texts.ProfileFillInfo, InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID()}

		c.services.Session.Set(query.From.ID, model.Session_FillProfileAge)

		_, err := bot.EditMessageText(&infoMsg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

		msg := tu.Message(chat.ChatID(), texts.AgeQuestion)

		_, err = bot.SendMessage(msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.CallbackDataEqual(FILL_PROFILE))

	// Age handler
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		const op = "handlers.callbacks.FillProfile.Message.Age"

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

		_, err := strconv.Atoi(textAge)
		if err != nil {
			ok = false
			msg = tu.Message(
				tu.ID(update.Message.Chat.ID),
				texts.AgeIncorrect,
			)
		}

		if ok {
			c.services.Session.Set(update.Message.From.ID, model.Session_FillProfileGender)
			c.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{Age: &age})
		}

		_, err = bot.SendMessage(msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.And(th.AnyMessageWithText(), predicates.ThMessageSessionEqual(*c.services, model.Session_FillProfileAge)))

	// Gender handler
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		const op = "handlers.callbacks.FillProfile.Message.Gender"

		chat := query.Message.GetChat()

		keyboard := tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton(texts.PhotoDefault).WithCallbackData(DEFAULT_PHOTO),
			),
		)

		msg := telego.EditMessageTextParams{Text: texts.PhotoInfo, InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID(), ReplyMarkup: keyboard}

		var gender string

		if query.Data == GENDER_MALE {
			gender = "M"
		}

		if query.Data == GENDER_FEMALE {
			gender = "F"
		}

		c.services.Session.Set(query.From.ID, model.Session_FillProfilePhoto)
		c.services.Profile.UpdateProfile(query.From.ID, userdto.ProfileUpdate{Gender: &gender})

		_, err := bot.EditMessageText(&msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.And(th.Or(th.CallbackDataEqual(GENDER_MALE), th.CallbackDataEqual(GENDER_FEMALE)), predicates.ThCallbackSessionEqual(*c.services, model.Session_FillProfileGender)))

	// Photo handler
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		const op = "handlers.callbacks.FillProfile.Message.Photo"

		var msg *telego.SendMessageParams

		photos, err := bot.GetUserProfilePhotos(&telego.GetUserProfilePhotosParams{UserID: query.From.ID, Limit: 1})
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

		ok := true

		if photos.TotalCount == 0 {
			ok = false
			msg = tu.Message(
				tu.ID(query.From.ID),
				"Не удалось получить ни одну фотографию, загрузите пожалуйста своё настоящее фото",
			)
		}

		if ok {
			fileId := photos.Photos[0][len(photos.Photos[0])-1].FileID
			c.log.Info(fileId)
		}

		_, err = bot.SendMessage(msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.And(th.CallbackDataEqual(DEFAULT_PHOTO), predicates.ThCallbackSessionEqual(*c.services, model.Session_FillProfilePhoto)))

	// About handler
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		const op = "handlers.callbacks.FillProfile.Message.About"

		about := update.Message.Text
		var msg *telego.SendMessageParams

		c.log.Info(about)

		keyboard := tu.InlineKeyboard(tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.GenderMale).WithCallbackData(GENDER_MALE),
			tu.InlineKeyboardButton(texts.GenderFemale).WithCallbackData(GENDER_FEMALE),
		))

		msg = tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Complete",
		).WithReplyMarkup(keyboard)

		c.services.Session.Set(update.Message.From.ID, model.Session_None)
		c.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{About: &about})

		_, err := bot.SendMessage(msg)
		if err != nil {
			c.log.Error(op, zap.Error(err))
		}

	}, th.And(th.AnyMessageWithText(), predicates.ThMessageSessionEqual(*c.services, model.Session_FillProfileAbout)))

}
