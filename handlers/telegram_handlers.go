package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	TelegramBotToken = "TOKEN"
	ChatID = "ID"
)

func SendTelegramMessageClient(name, phoneNumber, birthDate, adres string) error {
	bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(ChatID, "Добавлен новый клиент:\n\n"+
		"Имя: "+name+"\n"+
		"Телефон: "+phoneNumber+"\n"+
		"Дата рождения: "+birthDate+"\n"+
		"Адрес: "+adres)

	_, err = bot.Send(message)
	return err
}

func SendTelegramMessageSub(name, price, soldDate, clientID string) error {
	bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(ChatID, "Новая покупка:\n\n"+
		"Имя: "+name+"\n"+
		"Цена: "+price+"\n"+
		"Дата покупки: "+soldDate+"\n"+
		"Клиент: "+clientID)

	_, err = bot.Send(message)
	return err
}
