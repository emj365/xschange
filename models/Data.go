package models

import (
	"github.com/gorilla/websocket"
)

type websocketConnects map[*websocket.Conn]bool

type data struct {
	Orders  OrderList
	Users   UserList
	Clients map[string]websocketConnects
}

var Data = data{
	Orders: OrderList{},
	Users: UserList{
		&User{Balance: 100, GoodAmount: 100},
		&User{Balance: 100, GoodAmount: 100},
	},
	Clients: map[string]websocketConnects{
		"orders": make(websocketConnects),
		"matchs": make(websocketConnects),
	},
}
