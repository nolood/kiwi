package profile

import (
	"errors"
	"kiwi/.gen/kiwi/public/model"
	userdto "kiwi/internal/app/bot/dto/user"
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/static/texts"
	"kiwi/internal/app/bot/utils/predicates"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

func (s *Scene) handleAge(next func(chatId telego.ChatID)) {
	const op = "bot.handlers.scenes.profile.handleAge"

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
			err = s.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{Age: &age})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
			err = s.services.Session.Set(update.Message.From.ID, model.Session_None)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
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
	const op = "bot.handlers.scenes.profile.handleGender"

	genderMap := map[string]string{
		callbacks_consts.GENDER_MALE:   "M",
		callbacks_consts.GENDER_FEMALE: "F",
	}

	s.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		chat := query.Message.GetChat()

		gender, exists := genderMap[query.Data]
		if !exists {
			s.log.Error(op, zap.Error(errors.New("gender not found")))
		}

		err := s.services.Session.Set(query.From.ID, model.Session_None)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		err = s.services.Profile.UpdateProfile(query.From.ID, userdto.ProfileUpdate{Gender: &gender})

		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		next(chat.ChatID())

	}, th.And(th.Or(th.CallbackDataEqual(callbacks_consts.GENDER_MALE), th.CallbackDataEqual(callbacks_consts.GENDER_FEMALE)), predicates.ThCallbackSessionEqual(*s.services, model.Session_FillProfileGender)))
}

func (s *Scene) handleAbout(next func(chatId telego.ChatID)) {
	const op = "bot.handlers.scenes.profile.handleAbout"

	s.bh.Handle(func(bot *telego.Bot, update telego.Update) {

		about := update.Message.Text
		msg := tu.Message(
			tu.ID(update.Message.Chat.ID),
			texts.AboutComplete,
		)

		ok := true

		if len(about) > 300 {
			ok = false

			msg = tu.Message(
				tu.ID(update.Message.Chat.ID),
				texts.AboutLongText(len(about), 300),
			)
		}

		if ok {
			err := s.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{About: &about})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
			err = s.services.Session.Set(update.Message.From.ID, model.Session_None)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
		}

		_, err := bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			next(update.Message.Chat.ChatID())
		}

	}, th.And(th.AnyMessage(), predicates.ThMessageSessionEqual(*s.services, model.Session_FillProfileAbout)))
}

func (s *Scene) handlePhoto(next func(chatId telego.ChatID)) {
	const op = "bot.handlers.scenes.profile.handlePhoto"
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
			err := s.services.Profile.UpdateProfile(update.Message.From.ID, userdto.ProfileUpdate{PhotoId: &fileId})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
		}

		_, err := bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			chat := update.Message.GetChat()
			err = s.services.Session.Set(update.Message.From.ID, model.Session_None)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
			next(chat.ChatID())
		}

	}, th.And(th.AnyMessage(), predicates.ThMessageSessionEqual(*s.services, model.Session_FillProfilePhoto)))

}

func (s *Scene) handleDefaultPhoto(next func(chatId telego.ChatID)) {
	const op = "bot.handlers.scenes.profile.handleDefaultPhoto"

	s.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

		msg := tu.Message(
			tu.ID(query.From.ID),
			texts.PhotoComplete,
		)
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
			err = s.services.Profile.UpdateProfile(query.From.ID, userdto.ProfileUpdate{PhotoId: &fileId})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
			err = s.services.Session.Set(query.From.ID, model.Session_None)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
		}

		_, err = s.bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			next(chat.ChatID())
		}

	}, th.And(th.CallbackDataEqual(callbacks_consts.DEFAULT_PHOTO), predicates.ThCallbackSessionEqual(*s.services, model.Session_FillProfilePhoto)))
}

func (s *Scene) handleLocation(next func(chatId telego.ChatID)) {
	const op = "bot.handlers.scenes.profile.handleLocation"
	s.bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		chatId := message.From.ID

		if message.Location != nil {
			err := s.services.Profile.UpdateProfile(chatId, userdto.ProfileUpdate{Longitude: &message.Location.Longitude, Latitude: &message.Location.Latitude})
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}
		}

		data, err := s.services.MApp.Services.Location.Search(message.Text)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		kRows := make([][]telego.InlineKeyboardButton, 0)

		for _, hit := range data.Hits {
			s.log.Info(hit.GeneralMatch)
			s.log.Info("kek", zap.Any("hit", hit.MatchesPosition))

			kRows = append(kRows, tu.InlineKeyboardRow(
				tu.InlineKeyboardButton(hit.Name+"1").WithCallbackData("kek"),
			))
		}

		keyboard := tu.InlineKeyboard(kRows...)

		msg := tu.Message(message.Chat.ChatID(), "Выберите город").WithReplyMarkup(keyboard)

		_, err = s.bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

	}, th.And(th.AnyMessage(), predicates.ThMessageSessionEqual(*s.services, model.Session_FillProfileLocation)))
}
