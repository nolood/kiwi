package profile

import (
	"errors"
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	userdto "kiwi/internal/app/bot/dto/user"
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/static/texts"
	"kiwi/internal/app/bot/utils/predicates"
	"strconv"
	"strings"

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

		if len(data.Hits) == 0 {
			msg := tu.Message(message.Chat.ChatID(), texts.LocationNotFound)
			_, err = s.bot.SendMessage(msg)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}

			return
		}

		kRows := make([][]telego.InlineKeyboardButton, 0)

		for _, hit := range data.Hits {
			name := processLocationName(hit.Alternatenames, hit.MatchesPosition.Alternatenames)

			callback := callbacks_consts.LOCATION_TOWN + strconv.Itoa(int(hit.ID))

			kRows = append(kRows, tu.InlineKeyboardRow(
				tu.InlineKeyboardButton(name+" - "+hit.Name).WithCallbackData(callback),
			))

		}

		kRows = append(kRows, tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.LocationNotInList).WithCallbackData(callbacks_consts.LOCATION_TOWN+"not"),
		))

		keyboard := tu.InlineKeyboard(kRows...)

		msg := tu.Message(message.Chat.ChatID(), texts.LocationChoice).WithReplyMarkup(keyboard)

		_, err = s.bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

	}, th.And(th.AnyMessage(), predicates.ThMessageSessionEqual(*s.services, model.Session_FillProfileLocation)))
}

func (s *Scene) handleLocationTown(next func(chatId telego.ChatID)) {
	const op = "bot.handlers.scenes.profile.handleLocationTown"
	s.bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

		fmt.Println("HANDLE TOWN")

		chatId := tu.ID(query.From.ID)

		town := strings.Split(query.Data, "_")[1]

		var msg *telego.SendMessageParams

		ok := true

		if town == "not" {
			msg = tu.Message(
				chatId,
				texts.LocationNotFound,
			)

			ok = false
		}

		if ok {
			townId, err := strconv.Atoi(town)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}

			city, err := s.services.Cities.GetById(townId)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}

			s.services.Profile.UpdateProfile(query.From.ID, userdto.ProfileUpdate{Latitude: city.Latitude, Longitude: city.Longitude})

			err = s.services.Session.Set(query.From.ID, model.Session_None)
			if err != nil {
				s.log.Error(op, zap.Error(err))
			}

			msg = tu.Message(chatId, texts.LocationComplete)
		}

		_, err := s.bot.SendMessage(msg)
		if err != nil {
			s.log.Error(op, zap.Error(err))
		}

		if ok {
			next(chatId)
		}
	}, th.And(th.CallbackDataPrefix(callbacks_consts.LOCATION_TOWN), predicates.ThCallbackSessionEqual(*s.services, model.Session_FillProfileLocation)))
}
