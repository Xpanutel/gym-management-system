package models;

import "time"

type Client struct {
	ID 			int 
	Name	 	string
	BirthDate 	time.Time
	PhoneNumber string
	Adres 		string
	Subscriptions []Subscription
}