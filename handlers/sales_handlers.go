package handlers

import (
	"athleticclub/database"
	"athleticclub/models"
	"html/template"
	"net/http"
	"time"
	"log"
)

func PurchaseMembership(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db := database.GetDB()

		clientName := r.FormValue("client")
		subscriptionName := r.FormValue("subscription")
		employeeName := r.FormValue("employee")
		payment := r.FormValue("payment")

		purchase_date := time.Now()

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
		_, err = db.Exec("INSERT INTO sales (employee_id, client_id, subscription_id, price, payment, purchase_date) VALUES (?, ?, ?, ?, ?, ?)",
			employeeID, clientID, subscriptionID, price, payment, purchase_date)
		if err != nil {
			http.Error(w, "Ошибка при записи в базу данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = SendTelegramMessageSale(clientName, clientID, subscriptionName, price, employeeName, purchase_date)
		if err != nil {
			log.Println("Ошибка при отправке сообщения в Telegram:", err)
		}

		http.Redirect(w, r, "/sales", http.StatusSeeOther)
	}
}

func ShowPurchaseForm(w http.ResponseWriter, r *http.Request) {
	clients, err := GetClients() 
	if err != nil {
		http.Error(w, "Ошибка получения клиентов", http.StatusInternalServerError)
		return
	}

	subscriptions, err := GetSubs() 
	if err != nil {
		http.Error(w, "Ошибка получения клиентов", http.StatusInternalServerError)
		return
	}

	data := struct {
		Clients []models.Client
		Subscriptions []models.Subscription
	}{
		Clients:	clients,
		Subscriptions: subscriptions,
	}

	tmpl := template.Must(template.ParseFiles("templates/sales.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}