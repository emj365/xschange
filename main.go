package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type orderType bool
type orderStatus string

const (
	selling orderType   = false
	buying  orderType   = true
	open    orderStatus = "OPEN"
	closed  orderStatus = "CLOSE"
)

var sellings []*order
var buyings []*order
var orders []*order

// User data
type user struct {
	ID      uint
	Balance int
}

// Order created by the user
type order struct {
	UserID   uint
	Type     orderType
	Quantity int
	Price    int
	Matchs   []*match
	Status   orderStatus
}

// Match is one of trades of order
type match struct {
	Order    *order
	Quantity int
	Price    int
}

func (o *order) doSettlement() {
}

func (o *order) match() bool {
	return true
}

func (o *order) place() {
}

func (o *order) create() {
	orders = append(orders, o)
	if remain := o.match(); remain {
		o.place()
	} else {
		o.Status = closed
	}

	if len(o.Matchs) > 0 {
		o.doSettlement()
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders", postOrders).Methods("POST")
}

func extractOrderFromRequest(r *http.Request, o *order) {
	json.NewDecoder(r.Body).Decode(o)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func postOrders(w http.ResponseWriter, r *http.Request) {
	var o order
	extractOrderFromRequest(r, &o)
	o.create()
}
