package models

import "time"

type Subscription struct {
	ID			int
	Name		string
	Price		float64
	SoldDate	time.Time
	ClientID	int
}