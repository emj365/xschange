package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emj365/xschange/controllers"
	"github.com/emj365/xschange/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8000", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func createTestUsers() {
	models.Data.Users = append(models.Data.Users, &models.User{Balance: 0, GoodAmount: 0})
	models.Data.Users = append(models.Data.Users, &models.User{Balance: 0, GoodAmount: 0})
	for _, user := range models.Data.Users {
		models.Data.BalanceChanges.Add(&models.BalanceChange{
			Balance:           100,
			Good:              100,
			User:              user,
			IsDepositWithdrew: true,
		})
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	createTestUsers()

	router := mux.NewRouter()
	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders", controllers.PostOrders).Methods("POST")
	router.HandleFunc("/users", controllers.PostUsers).Methods("POST")
	router.HandleFunc("/subscriptions/orders", subscribeOrders)
	router.HandleFunc("/subscriptions/matchs", subscribeMatchs)

	log.Printf("server is running on %v", *addr)
	log.Fatal(http.ListenAndServe(*addr, router))
}

func subscribeOrders(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register client
	models.Data.Clients["orders"][ws] = true
}

func subscribeMatchs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register client
	models.Data.Clients["matchs"][ws] = true
}
