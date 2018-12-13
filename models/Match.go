package models

// Match is one of trades of order
type Match struct {
	Order    *Order
	Quantity uint
	Price    uint
}
