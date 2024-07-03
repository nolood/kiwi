package handlers

import (
	"kiwi/internal/app/bot/handlers/callbacks"
	"kiwi/internal/app/bot/handlers/commands"
	"kiwi/internal/app/bot/services"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"

	"go.uber.org/zap"
)

func Register(log *zap.Logger, updates <-chan telego.Update, b *telego.Bot, servs *services.Services) {
	bh, _ := th.NewBotHandler(b, updates)

	comms := commands.New(log, servs)
	callbacks := callbacks.New(log, servs)

	callbacks.ViewProfile(bh)
	comms.Start(bh)

	bh.Start()
}
