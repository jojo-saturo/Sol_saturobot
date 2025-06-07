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
		msg, err := utils.CreateSolAccount()
		if err != nil {
			return c.Send("❌ Failed to create Solana account.")
		}
		return c.Send(msg)
	})

	bot.Start()
	return nil
}
