package models

// Match is one of trades of order
type Match struct {
	Order    *Order
	Quantity int
	Price    int
}
