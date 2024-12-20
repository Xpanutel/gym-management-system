package handlers

import (
	"database/sql"
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
	"net/http"
)

// создание нового клиента
func addClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db := database.GetDB()
		name := r.FormValue("name")
		phoneNumber := r.FormValue("phoneNumber")
		birthDate := r.FormValue("birthDate")
		adres := r.FormValue("adres")

		_, err := db.Exec("INSERT INTO clients(name, birtch_date, phone_nubmer, adres)
		VALUES (?,?,?);")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w,r, "/clients", http.StatusSeeOther)
	}
}

// отображение страницы + клиентов
func showClients (w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close();

	var clients []model.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.BirthDate, &client.PhoneNumber, &client.Adres); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
		clients = append(clients, client)
	}

	tmpl := template.Must(template.ParseFiles("templates/cients.html"))
	tmpl.Execute(w, clients)
}

