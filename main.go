package main

type User struct {
	ID      uint
	balance int
}

type Order struct {
	Quantity int
	Price    int
	UserID   uint
	Sell     bool
}

var sellingOrders []Order
var buyingOrders []Order
var orders []Order
