package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	TelegramBotToken = "7375756872:AAG6KJ9JnldMypthj2Rgv-b-HPPaM10sMjA"
	ChatID = 1456988449
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

func SendTelegramMessageSub(name, price, soldDate, clientID, employeeID string) error {
	bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(ChatID, "Новая покупка:\n\n"+
		"Имя: "+name+"\n"+
		"Цена: "+price+"\n"+
		"Дата покупки: "+soldDate+"\n"+
		"Клиент: "+clientID+"\n"+
		"Сотрудник: "+employeeID)

	_, err = bot.Send(message)
	return err
}