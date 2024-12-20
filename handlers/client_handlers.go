package handlers

import (
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
	"net/http"
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