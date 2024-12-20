package handlers

import (
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
	"net/http"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	TelegramBotToken = "TOKEN"
	ChatID = "CHATID"
)

func sendTelegramMessage(name, phoneNumber, birthDate, adres string) error {
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

// Создание нового клиента
func AddClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db := database.GetDB()
		name := r.FormValue("name")
		phoneNumber := r.FormValue("phoneNumber")
		birthDate := r.FormValue("birthDate")
		adres := r.FormValue("adres")
		
		_, err := db.Exec("INSERT INTO clients(name, birth_date, phone_number, adres) VALUES (?,?,?,?);", 
		name, birthDate, phoneNumber, adres)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Отправка сообщения в Telegram
		err = sendTelegramMessage(name, phoneNumber, birthDate, adres)
		if err != nil {
			log.Println("Ошибка при отправке сообщения в Telegram:", err)
		}

		http.Redirect(w, r, "/clients", http.StatusSeeOther)
	}
}

// Отображение страницы клиентов
func ShowClients(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.BirthDate, &client.PhoneNumber, &client.Adres); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
		clients = append(clients, client)
	}

	tmpl := template.Must(template.ParseFiles("templates/clients.html"))
	tmpl.Execute(w, clients)
}