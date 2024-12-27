package handlers

import (
	"athleticclub/database"
	"athleticclub/models"
	"html/template"
	"net/http"
	"time"
	"log"
	"fmt"
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

func GetSalesReport(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Получаем дату из запроса
		dateStr := r.URL.Query().Get("date")
		if dateStr == "" {
			http.Error(w, "Дата не указана", http.StatusBadRequest)
			return
		}

		// Преобразуем строку даты в формат time.Time
		purchaseDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Некорректный формат даты", http.StatusBadRequest)
			return
		}

		db := database.GetDB()
		rows, err := db.Query(`
			SELECT c.name AS client_name, e.name AS employee_name, s.name AS subscription_name, sale.price, sale.purchase_date
			FROM sales AS sale
			JOIN clients AS c ON sale.client_id = c.id
			JOIN employees AS e ON sale.employee_id = e.id
			JOIN subscriptions AS s ON sale.subscription_id = s.id
			WHERE DATE(sale.purchase_date) = ?`, purchaseDate.Format("2006-01-02"))
		if err != nil {
			http.Error(w, "Ошибка при получении данных из базы: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Формируем отчет
		var report []string
		for rows.Next() {
			var clientName, employeeName, subscriptionName string
			var price float64
			var purchaseDateRaw []byte 

			err := rows.Scan(&clientName, &employeeName, &subscriptionName, &price, &purchaseDateRaw)
			if err != nil {
				http.Error(w, "Ошибка при чтении данных: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Преобразуем []byte в time.Time
			purchaseDate, err := time.Parse("2006-01-02 15:04:05", string(purchaseDateRaw))
			if err != nil {
				http.Error(w, "Ошибка при преобразовании даты: "+err.Error(), http.StatusInternalServerError)
				return
			}

			report = append(report, fmt.Sprintf("Клиент: %s, Продавец: %s, Продукт: %s, Цена: %.2f, Дата: %s", 
				clientName, employeeName, subscriptionName, price, purchaseDate.Format("2006-01-02 15:04:05")))
		}

		// Проверяем, есть ли продажи за указанную дату
		if len(report) == 0 {
			fmt.Fprintln(w, "Нет продаж за указанную дату.")
			return
		}

		// Выводим отчет
		for _, entry := range report {
			fmt.Fprintln(w, entry)
		}
	}
}
