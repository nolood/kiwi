package commands

import (
	callbacks_consts "kiwi/internal/app/bot/handlers/callbacks/consts"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/app/bot/static/texts"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.uber.org/zap"
)

const (
	START = "start"
)

type Commands struct {
	log      *zap.Logger
	services *services.Services
	Bh       *th.BotHandler
}

func New(log *zap.Logger, servs *services.Services, b *telego.Bot, updates <-chan telego.Update) *Commands {

	bh, err := th.NewBotHandler(b, updates)
	if err != nil {
		log.Fatal("handlers.commands.New", zap.Error(err))
	}

	return &Commands{
		log:      log,
		services: servs,
		Bh:       bh,
	}
}

func (c *Commands) Start() {
	c.Bh.Handle(func(bot *telego.Bot, update telego.Update) {

		userprof, err := c.services.User.GetOrCreate(update.Message.From)
		if err != nil {
			c.log.Error("handlers.commands.Start", zap.Error(err))
			return
		}

		var keyboard *telego.InlineKeyboardMarkup

		// Если есть заполненная анкета и отключена - | Начать поиск | Посмотреть анкету |
		if userprof.Profile.Age != nil && userprof.Profile.Gender != nil && !userprof.Profile.IsActive {
			keyboard = tu.InlineKeyboard(tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Начать поиск").WithCallbackData("search"),
				tu.InlineKeyboardButton("Посмотреть анкету").WithCallbackData(callbacks_consts.VIEW_PROFILE),
			))
		}

		// Если анкета заполнена не до конца | Заполнить анкету |

		if userprof.Profile.Age == nil || userprof.Profile.Gender == nil {
			keyboard = tu.InlineKeyboard(tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Заполнить анкету").WithCallbackData(callbacks_consts.FILL_PROFILE),
			))
		}

		// Если анкета заполнена и включена - | Продолжить поиск | Посмотреть анкету |

		if userprof.Profile.Age != nil && userprof.Profile.Gender != nil && userprof.Profile.IsActive {
			keyboard = tu.InlineKeyboard(tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Продолжить поиск").WithCallbackData("start"),
				tu.InlineKeyboardButton("Посмотреть анкету").WithCallbackData(callbacks_consts.VIEW_PROFILE),
			))
		}

		msg := tu.Message(
			tu.ID(update.Message.Chat.ID),
			texts.GreetingText(update.Message.From.FirstName),
		).WithReplyMarkup(keyboard)

		_, _ = bot.SendMessage(msg)

	}, th.CommandEqual(START))
}
