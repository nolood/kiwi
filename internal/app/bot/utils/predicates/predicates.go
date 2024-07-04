package predicates

import (
	"kiwi/.gen/kiwi/public/model"
	"kiwi/internal/app/bot/services"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

func ThMessageSessionEqual(servs services.Services, session model.Session) telegohandler.Predicate {
	return func(update telego.Update) bool {

		if update.Message == nil {
			return false
		}

		return servs.Session.Compare(update.Message.From.ID, session)
	}
}

func ThCallbackSessionEqual(servs services.Services, session model.Session) telegohandler.Predicate {
	return func(update telego.Update) bool {

		if update.CallbackQuery == nil {
			return false
		}

		return servs.Session.Compare(update.CallbackQuery.From.ID, session)
	}
}
