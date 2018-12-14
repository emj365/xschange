package main

import (
	"log"
	"net/http"

	"github.com/emj365/xschange/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders", controllers.PostOrders).Methods("POST")
	router.HandleFunc("/users", controllers.PostUsers).Methods("POST")
	log.Println("server is running on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
