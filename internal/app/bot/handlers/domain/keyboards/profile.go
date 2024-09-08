package keyboards

import (
	"kiwi/internal/app/bot/static/texts"

	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ProfileEditKeyboard(rows ...[]telego.InlineKeyboardButton) *telego.InlineKeyboardMarkup {

	keyboardRows := [][]telego.InlineKeyboardButton{
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileFillAgain).WithCallbackData(callbacks_consts.EDIT_PROFILE + "again"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileEditName).WithCallbackData(callbacks_consts.EDIT_PROFILE + "name"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileEditAge).WithCallbackData(callbacks_consts.EDIT_PROFILE + "age"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileEditGender).WithCallbackData(callbacks_consts.EDIT_PROFILE + "gender"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileEditPhoto).WithCallbackData(callbacks_consts.EDIT_PROFILE + "photo"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileEditAbout).WithCallbackData(callbacks_consts.EDIT_PROFILE + "about"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(texts.ProfileEditLocation).WithCallbackData(callbacks_consts.EDIT_PROFILE + "location"),
		),
	}

	keyboardRows = append(keyboardRows, rows...)

	return tu.InlineKeyboard(keyboardRows...)
}
