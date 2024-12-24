package handlers

import (
	"net/http"
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
)

func AddSub(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		database := database.GetDB()

		name := r.FormValue("name")
		price := r.FormValue("price")
		period := r.FormValue("period")

		_, err := database.Exec("INSERT INTO subscriptions (name, price, period) VALUES (?,?,?);", name, price, period)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/subs", http.StatusSeeOther)
	}
}

func ShowSubs(w http.ResponseWriter, r *http.Request) {
	database := database.GetDB()

	rows, err := database.Query("SELECT * FROM subscriptions")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(&subscription.ID, &subscription.Name, &subscription.Price, &subscription.Period); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		subscriptions = append(subscriptions, subscription)
	}

	tmpl := template.Must(template.ParseFiles("templates/subs.html"))
	if err := tmpl.Execute(w, subscriptions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}