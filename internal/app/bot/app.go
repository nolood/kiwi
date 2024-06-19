package bot

import (
	"fmt"
	"kiwi/internal/config"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type BotApp struct {
	cfg config.Telegram
	Bot *telego.Bot
}

func New(cfg config.Telegram) *BotApp {
	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &BotApp{
		cfg: cfg,
		Bot: bot,
	}
}

func (b *BotApp) MustRun() {
	updates, _ := b.Bot.UpdatesViaLongPolling(nil)
	defer b.Bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			sentMessage, _ := b.Bot.SendMessage(
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}

}

func (b *BotApp) Stop() {
	b.Bot.StopLongPolling()
}
