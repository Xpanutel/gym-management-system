package models

type Subscription struct {
	ID			int
	Name		string
	Price		float64
	SoldDate	time.Time
	ClientID	int
}