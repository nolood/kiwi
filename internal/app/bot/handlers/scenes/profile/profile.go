package profile

import (
	"kiwi/.gen/kiwi/public/model"
	userdto "kiwi/internal/app/bot/dto/user"
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/bot/static/texts"
	"kiwi/internal/app/bot/utils/predicates"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

type Scene struct {
	log      *zap.Logger
	services *services.Services
	Bh       *th.BotHandler
	bot      *telego.Bot
}

func New(log *zap.Logger, servs *services.Services, bot *telego.Bot, updates <-chan telego.Update) *Scene {

	const op = "handlers.scenes.New"

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatal(op, zap.Error(err))
	}

	return &Scene{
		log:      log,
		services: servs,
		Bh:       bh,
		bot:      bot,
	}
}

func (s *Scene) handleAge(next func(chatId telego.ChatID)) {
	const op = "handlers.scenes.profile.handleAge"

	s.log.Debug("Register profile age handler")

	s.Bh.Handle(func(bot *telego.Bot, update telego.Update) {

		s.log.Info("AGE HANDLER triggered")

		textAge := update.Message.Text
		var msg *telego.SendMessageParams

		age, _ := strconv.Atoi(textAge)

		ok := true

		msg = tu.Message(
			tu.ID(update.Message.Chat.ID),
			texts.AgeComplete,
		)

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
			s.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{Age: &age})
			s.services.Session.Set(update.Message.From.ID, model.Session_None)
		}

		_, err = bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			next(update.Message.Chat.ChatID())
		}

	}, th.And(th.AnyMessageWithText(), predicates.ThMessageSessionEqual(*s.services, model.Session_FillProfileAge)))

}

func (s *Scene) handleGender(next func(chatId telego.ChatID)) {
	s.Bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		const op = "handlers.scenes.profile.handleGender"

		chat := query.Message.GetChat()

		keyboard := tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton(texts.PhotoDefault).WithCallbackData(callbacks_consts.DEFAULT_PHOTO),
			),
		)

		msg := telego.EditMessageTextParams{Text: texts.PhotoInfo, InlineMessageID: query.InlineMessageID, MessageID: query.Message.GetMessageID(), ChatID: chat.ChatID(), ReplyMarkup: keyboard}

		var gender string

		if query.Data == callbacks_consts.GENDER_MALE {
			gender = "M"
		}

		if query.Data == callbacks_consts.GENDER_FEMALE {
			gender = "F"
		}

		s.services.Session.Set(query.From.ID, model.Session_FillProfilePhoto)
		s.services.Profile.UpdateProfile(query.From.ID, userdto.ProfileUpdate{Gender: &gender})

		_, err := bot.EditMessageText(&msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		next(chat.ChatID())

	}, th.And(th.Or(th.CallbackDataEqual(callbacks_consts.GENDER_MALE), th.CallbackDataEqual(callbacks_consts.GENDER_FEMALE)), predicates.ThCallbackSessionEqual(*s.services, model.Session_FillProfileGender)))
}

func (s *Scene) GetAge(chatId telego.ChatID) {
	const op = "handlers.scenes.profile.GetAge"

	s.services.Session.Set(chatId.ID, model.Session_FillProfileAge)

	msg := tu.Message(chatId, texts.AgeQuestion)

	_, err := s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

}

// GetGender - get gender from user input
func (s *Scene) GetGender(chatId telego.ChatID) {
	const op = "handlers.scenes.profile.GetGender"

	s.services.Session.Set(chatId.ID, model.Session_FillProfileGender)

	keyboard := tu.InlineKeyboard(tu.InlineKeyboardRow(
		tu.InlineKeyboardButton(texts.GenderMale).WithCallbackData(callbacks_consts.GENDER_MALE),
		tu.InlineKeyboardButton(texts.GenderFemale).WithCallbackData(callbacks_consts.GENDER_FEMALE),
	))

	msg := tu.Message(
		chatId,
		texts.GenderQuestion,
	).WithReplyMarkup(keyboard)

	_, err := s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) FillProfile(chatId telego.ChatID) {

	s.handleAge(func(chatId telego.ChatID) {
		s.GetGender(chatId)
	})

	s.handleGender(func(chatId telego.ChatID) {
		s.log.Info("FILL PROFILE COMPLETE)))")
	})

	s.GetAge(chatId)

	s.log.Debug("Start profile bot handler")
	s.Bh.Start()
}
