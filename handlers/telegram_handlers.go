package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
	"fmt"
	"strconv"
)

const (
	TelegramBotToken = "TOKEN"
	ChatID = CHATID
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

func SendTelegramMessageSale(clientName string, clientID int, subscriptionName string, price float64, employeeName string, purchase_date time.Time) error {
	bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(ChatID, "🔔 Новая покупка зарегистрирована 🔔\n\n"+
		"🔹 Имя покупателя: "+clientName+"\n"+
		"🔹 ID клиента: "+strconv.Itoa(clientID)+"\n"+
		"🔹 Имя покупателя: "+subscriptionName+"\n"+
		"🔹 Сумма покупки: "+fmt.Sprintf("%.2f", price)+" ₽\n"+
		"🔹 Дата и время: "+purchase_date.Format("2006-01-02 15:04:05")+"\n"+
		"🔹 Ответственный сотрудник: "+employeeName)

	_, err = bot.Send(message)
	return err
}
