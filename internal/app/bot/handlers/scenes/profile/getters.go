package profile

import (
	"kiwi/.gen/kiwi/public/model"
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/static/texts"

	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

func (s *Scene) GetAge(chatId telego.ChatID) {
	const op = "bot.handlers.scenes.profile.GetAge"

	err := s.services.Session.Set(chatId.ID, model.Session_FillProfileAge)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	msg := tu.Message(chatId, texts.AgeQuestion)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

}

func (s *Scene) GetGender(chatId telego.ChatID) {
	const op = "bot.handlers.scenes.profile.GetGender"

	err := s.services.Session.Set(chatId.ID, model.Session_FillProfileGender)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	keyboard := tu.InlineKeyboard(tu.InlineKeyboardRow(
		tu.InlineKeyboardButton(texts.GenderMale).WithCallbackData(callbacks_consts.GENDER_MALE),
		tu.InlineKeyboardButton(texts.GenderFemale).WithCallbackData(callbacks_consts.GENDER_FEMALE),
	))

	msg := tu.Message(
		chatId,
		texts.GenderQuestion,
	).WithReplyMarkup(keyboard)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) GetPhoto(chatId telego.ChatID) {
	const op = "bot.handlers.scenes.profile.GetPhoto"

	err := s.services.Session.Set(chatId.ID, model.Session_FillProfilePhoto)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.PhotoDefault).WithCallbackData(callbacks_consts.DEFAULT_PHOTO),
		),
	)

	msg := tu.Message(chatId, texts.PhotoInfo).WithReplyMarkup(keyboard)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) GetAbout(chatId telego.ChatID) {
	const op = "bot.handlers.scenes.profile.GetAbout"
	err := s.services.Session.Set(chatId.ID, model.Session_FillProfileAbout)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
	msg := tu.Message(chatId, texts.AboutQuestion)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) StartFillProfileScene(chatId telego.ChatID) {
	s.GetAge(chatId)
}
