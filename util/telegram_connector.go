package util

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

import "log"

func TeleConnect(token string, messages chan string, telegramChannel int64) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for {
		select {
		case msg := <-messages:
			{
				if msg == "" {
					log.Printf("Получено пустое сообщение")
					continue
				}
				telemsg := tgbotapi.NewMessage(telegramChannel, msg)
				_, err := bot.Send(telemsg)
				if err != nil {
					log.Printf("Ошибка при отправке сообщения: %s", err)
				}
			}
		}
	}
	/*	u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	*/
	//return *bot, err
}
