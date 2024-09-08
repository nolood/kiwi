package profile

import (
	"kiwi/.gen/kiwi/public/model"
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/handlers/domain/keyboards"
	"kiwi/internal/app/bot/static/texts"

	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

func (s *Scene) GetName(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetName"

	var _session model.Session

	if session == nil {
		_session = model.Session_FillProfileName
	} else {
		_session = session.(model.Session)
	}

	err := s.services.Session.Set(chatId.ID, _session)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	msg := tu.Message(chatId, texts.NameQuestion)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) GetAge(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetAge"

	var _session model.Session

	if session == nil {
		_session = model.Session_FillProfileAge
	} else {
		_session = session.(model.Session)
	}

	err := s.services.Session.Set(chatId.ID, _session)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	msg := tu.Message(chatId, texts.AgeQuestion)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) GetLocation(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetLocation"

	var _session model.Session

	if session == nil {
		_session = model.Session_FillProfileLocation
	} else {
		_session = session.(model.Session)
	}

	err := s.services.Session.Set(chatId.ID, _session)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	keyboard := tu.Keyboard(
		tu.KeyboardRow(tu.KeyboardButton(texts.LocationSend).WithRequestLocation()),
	).WithResizeKeyboard().WithInputFieldPlaceholder(texts.LocationTown)

	msg := tu.Message(chatId, texts.LocationQuestion).WithReplyMarkup(keyboard)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) GetGender(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetGender"

	var _session model.Session

	if session == nil {
		_session = model.Session_FillProfileGender
	} else {
		_session = session.(model.Session)
	}

	err := s.services.Session.Set(chatId.ID, _session)
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

func (s *Scene) GetPhoto(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetPhoto"

	var _session model.Session

	if session == nil {
		_session = model.Session_FillProfilePhoto
	} else {
		_session = session.(model.Session)
	}

	err := s.services.Session.Set(chatId.ID, _session)
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

func (s *Scene) GetAbout(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetAbout"

	var _session model.Session

	if session == nil {
		_session = model.Session_FillProfileAbout
	} else {
		_session = session.(model.Session)
	}

	err := s.services.Session.Set(chatId.ID, _session)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
	msg := tu.Message(chatId, texts.AboutQuestion)

	_, err = s.bot.SendMessage(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}
}

func (s *Scene) GetProfileComplete(chatId telego.ChatID, session interface{}) {
	const op = "bot.handlers.scenes.profile.GetProfileComplete"

	err := s.services.Session.Set(chatId.ID, model.Session_EditProfile)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	msgKeyboardRemove := tu.Message(chatId, texts.ProfileFillComplete).WithReplyMarkup(tu.ReplyKeyboardRemove())

	_, err = s.bot.SendMessage(msgKeyboardRemove)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	inlineKeyboard := keyboards.ProfileEditKeyboard()

	msg, err := s.services.Profile.GetFormattedProfile(chatId)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

	msg.WithReplyMarkup(inlineKeyboard)

	_, err = s.bot.SendPhoto(msg)
	if err != nil {
		s.log.Error(op, zap.Error(err))
	}

}

func (s *Scene) StartFillProfileScene(chatId telego.ChatID) {
	s.GetName(chatId, nil)
}
