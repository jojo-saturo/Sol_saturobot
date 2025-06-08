package config

import (
	"fmt"
	"os"
	"time"

	"Saturobot/utils"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func StartTelegramBot() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	token := os.Getenv("TOKEN_KEY")
	if token == "" {
		return fmt.Errorf("TOKEN_KEY not found in environment")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	bot.Handle("/start", func(c tele.Context) error {
		msg, err := utils.StartSolBot()
		if err != nil {
			return c.Send("❌ Failed to Start Solana Bot.")
		}
		menu := &tele.ReplyMarkup{ResizeKeyboard: true}
		createBtn := menu.Text("🪙 Create Wallet")
		menu.Reply(menu.Row(createBtn))

		return c.Send(msg, menu)
	})
	bot.Handle(tele.OnText, func(c tele.Context) error {
		if c.Text() == "🪙 Create Wallet" {
			msg, err := utils.CreateSolAccount()
			if err != nil {
				return c.Send("❌ Failed to create wallet.")
			}
			return c.Send(msg)
		}
		return nil
	})

	bot.Handle("create", func(c tele.Context) error {
		msg, err := utils.CreateSolAccount()
		if err != nil {
			return c.Send("❌ Failed to create wallet.")
		}
		return c.Send(msg)
	})

	bot.Handle("/help", func(c tele.Context) error {
		return c.Send("ℹ️ Use /start to begin or tap 'Create Wallet' to generate a wallet.")
	})

	bot.Start()
	return nil
}
