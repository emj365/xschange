package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emj365/xschange/models"
	"github.com/gorilla/mux"
)

var orders = []*models.Order{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders", postOrders).Methods("POST")
	log.Println("server is running on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func extractOrderFromRequest(r *http.Request, o *models.Order) error {
	err := json.NewDecoder(r.Body).Decode(o)
	return err
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func postOrders(w http.ResponseWriter, r *http.Request) {
	o := models.Order{}
	err := extractOrderFromRequest(r, &o)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	o.Place(&orders)
	json.NewEncoder(w).Encode(o)
}
