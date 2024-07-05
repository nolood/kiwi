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
	bh       *th.BotHandler
	bot      *telego.Bot
}

func New(log *zap.Logger, servs *services.Services, bot *telego.Bot, bh *th.BotHandler) *Scene {

	return &Scene{
		log:      log,
		services: servs,
		bh:       bh,
		bot:      bot,
	}
}

func (s *Scene) handleAge(next func(chatId telego.ChatID)) {
	const op = "handlers.scenes.profile.handleAge"

	s.bh.Handle(func(bot *telego.Bot, update telego.Update) {
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
	s.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		chat := query.Message.GetChat()

		var gender string

		if query.Data == callbacks_consts.GENDER_MALE {
			gender = "M"
		}

		if query.Data == callbacks_consts.GENDER_FEMALE {
			gender = "F"
		}

		s.services.Session.Set(query.From.ID, model.Session_None)
		s.services.Profile.UpdateProfile(query.From.ID, userdto.ProfileUpdate{Gender: &gender})

		next(chat.ChatID())

	}, th.And(th.Or(th.CallbackDataEqual(callbacks_consts.GENDER_MALE), th.CallbackDataEqual(callbacks_consts.GENDER_FEMALE)), predicates.ThCallbackSessionEqual(*s.services, model.Session_FillProfileGender)))
}

func (s *Scene) handlePhoto(next func(chatId telego.ChatID)) {
	const op = "handlers.scenes.profile.handlePhoto"
	s.bh.Handle(func(bot *telego.Bot, update telego.Update) {
		msg := tu.Message(
			tu.ID(update.Message.Chat.ID),
			texts.PhotoComplete,
		)

		ok := true

		if update.Message.Photo == nil {
			ok = false

			msg = tu.Message(
				tu.ID(update.Message.Chat.ID),
				texts.PhotoEmpty,
			)
		}

		var fileId string

		if ok {
			photos := update.Message.Photo
			fileId = photos[len(photos)-1].FileID
			s.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{PhotoId: &fileId})
		}

		_, err := bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			chat := update.Message.GetChat()
			_, err = s.bot.SendPhoto(&telego.SendPhotoParams{Photo: telego.InputFile{
				FileID: fileId}, ChatID: chat.ChatID()})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
			s.services.Session.Set(update.Message.From.ID, model.Session_None)
			next(chat.ChatID())
		}

	}, th.And(th.AnyMessage(), predicates.ThMessageSessionEqual(*s.services, model.Session_FillProfilePhoto)))

}

func (s *Scene) handleDefaultPhoto(next func(chatId telego.ChatID)) {
	const op = "handlers.scenes.profile.handleDefaultPhoto"

	s.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

		s.log.Info("default photo trigger")

		var msg *telego.SendMessageParams
		chat := query.Message.GetChat()

		photos, err := s.bot.GetUserProfilePhotos(&telego.GetUserProfilePhotosParams{UserID: query.From.ID, Limit: 1})
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		ok := true

		if photos.TotalCount == 0 {
			ok = false
			msg = tu.Message(
				tu.ID(query.From.ID),
				texts.PhotoDefaultEmpty,
			)
		}

		var fileId string

		if ok {
			fileId = photos.Photos[0][len(photos.Photos[0])-1].FileID
			s.log.Info(fileId)
		}

		_, err = s.bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			_, err = s.bot.SendPhoto(&telego.SendPhotoParams{Photo: telego.InputFile{
				FileID: fileId}, ChatID: chat.ChatID()})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
			s.services.Session.Set(query.From.ID, model.Session_None)
			next(chat.ChatID())
		}

	}, th.And(th.CallbackDataEqual(callbacks_consts.DEFAULT_PHOTO), predicates.ThCallbackSessionEqual(*s.services, model.Session_FillProfilePhoto)))
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

func (s *Scene) GetPhoto(chatId telego.ChatID) {
	const op = "handlers.scenes.profile.GetPhoto"

	s.services.Session.Set(chatId.ID, model.Session_FillProfilePhoto)

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.PhotoDefault).WithCallbackData(callbacks_consts.DEFAULT_PHOTO),
		),
	)

	msg := tu.Message(chatId, texts.PhotoInfo).WithReplyMarkup(keyboard)

	_, err := s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) StartFillProfileScene(chatId telego.ChatID) {
	s.GetAge(chatId)
}

func (s *Scene) RegisterFillProfileScene() {
	s.handleAge(func(chatId telego.ChatID) {
		s.GetGender(chatId)
	})

	s.handleGender(func(chatId telego.ChatID) {
		s.GetPhoto(chatId)
	})

	s.handleDefaultPhoto(func(chatId telego.ChatID) {
		s.log.Info("photo complete")
	})

	s.handlePhoto(func(chatId telego.ChatID) {
		s.log.Info("photo complete")

	})
}
