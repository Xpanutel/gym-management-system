package handlers

import (
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
	"net/http"
	"log"
)

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
		err = SendTelegramMessageClient(name, phoneNumber, birthDate, adres)
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

func GetClients() ([]models.Client, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT id, name FROM clients;")
	if err != nil {
		log.Println("Ошибка выполнения запроса: ", err)
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name); err != nil {
			log.Println("Ошибка при сканировании строки: ", err)
			return nil, err
		}

		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка при переборе строк:", err)
		return nil, err
	}

	return clients, nil
}