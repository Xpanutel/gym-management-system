package models;

type Client struct {
	ID 			int 
	Name	 	string
	BirthDate 	string
	PhoneNumber string
	Adres 		string
	Subscriptions []Subscription
}