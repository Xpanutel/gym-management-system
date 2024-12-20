package handlers

import (
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
	"net/http"
	"log"
)

func AddSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db := database.GetDB()
		name := r.FormValue("name")
		price := r.FormValue("price")
		soldDate := r.FormValue("sold_date")
		clientID := r.FormValue("client_id")

		_, err := db.Exec("INSERT INTO subscriptions (name, price, sold_date, client_id) VALUES (?,?,?,?);", name, price, soldDate, clientID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = SendTelegramMessageSub(name, price, soldDate, clientID) 
		if err != nil {
			log.Println("Ошибка при отправке сообщения в Telegram:", err)
		}

		http.Redirect(w,r, "/subs", http.StatusSeeOther)
	}
}

func ShowSubscriptions(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM subscriptions")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(&subscription.ID, &subscription.Name, &subscription.Price, &subscription.SoldDate, &subscription.ClientID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
		subs = append(subs, subscription)
	}

	tmpl := template.Must(template.ParseFiles("templates/subs.html"))
	tmpl.Execute(w, subs)
}
