package handlers

import (
	"athleticclub/database"
	"athleticclub/models"
	"html/template"
	"net/http"
	"time"
)

func PurchaseMembership(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db := database.GetDB()

		clientName := r.FormValue("client")
		subscriptionName := r.FormValue("subscription")
		employeeName := r.FormValue("employee")

		// Получаем ID клиента
		var clientID int
		err := db.QueryRow("SELECT id FROM clients WHERE name = ?", clientName).Scan(&clientID)
		if err != nil {
			http.Error(w, "Клиент не найден", http.StatusBadRequest)
			return
		}

		// Получаем ID абонемента и цену
		var subscriptionID int
		var price float64
		err = db.QueryRow("SELECT id, price FROM subscriptions WHERE name = ?", subscriptionName).Scan(&subscriptionID, &price)
		if err != nil {
			http.Error(w, "Абонемент не найден", http.StatusBadRequest)
			return
		}

		// Получаем ID сотрудника
		var employeeID int
		err = db.QueryRow("SELECT id FROM employees WHERE name = ?", employeeName).Scan(&employeeID)
		if err != nil {
			http.Error(w, "Сотрудник не найден", http.StatusBadRequest)
			return
		}

		// Вставляем запись о покупке
		_, err = db.Exec("INSERT INTO sales (employee_id, client_id, subscription_id, price, purchase_date) VALUES (?, ?, ?, ?, ?)",
			employeeID, clientID, subscriptionID, price, time.Now())
		if err != nil {
			http.Error(w, "Ошибка при записи в базу данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/sales", http.StatusSeeOther)
	}
}

func ShowPurchaseForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/sales.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

