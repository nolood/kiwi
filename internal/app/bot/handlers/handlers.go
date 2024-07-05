package handlers

import (
	"kiwi/internal/app/bot/handlers/callbacks"
	"kiwi/internal/app/bot/handlers/commands"
	"kiwi/internal/app/bot/handlers/scenes"
	"kiwi/internal/app/bot/services"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"

	"go.uber.org/zap"
)

func Register(log *zap.Logger, updates <-chan telego.Update, b *telego.Bot, servs *services.Services) {

	bh, err := th.NewBotHandler(b, updates)
	if err != nil {
		log.Fatal("handlers.Register", zap.Error(err))
	}

	defer bh.Stop()

	sc := scenes.New(log, servs, b, bh)
	sc.Profile.RegisterFillProfileScene()

	comms := commands.New(log, servs, b, bh)
	comms.Register()

	cb := callbacks.New(log, servs, b, bh, sc)
	cb.Register()

	bh.Start()

}
