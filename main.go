package main 

import (
	"log"
	"net/http"
	"athleticclub/handlers"
	"athleticclub/database"
)

func main() {
	database.InitDB("root:@tcp(localhost:3306)/athletic")

	http.HandleFunc("/clients", handlers.ShowClients)
	http.HandleFunc("/add-client", handlers.AddClient)

	http.HandleFunc("/employees", handlers.ShowEmployees)
	http.HandleFunc("/add-emp", handlers.AddEmployee)

	http.HandleFunc("/subs", handlers.ShowSubs)
	http.HandleFunc("/add-sub", handlers.AddSub)
	
	http.HandleFunc("/add-sale", handlers.PurchaseMembership)
	http.HandleFunc("/sales", handlers.ShowPurchaseForm)
	http.HandleFunc("/sales/reports", handlers.GetSalesReport)


	log.Println("Сервер запущен на localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
