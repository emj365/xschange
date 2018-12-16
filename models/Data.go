package models

import (
	"github.com/gorilla/websocket"
)

type websocketConnects map[*websocket.Conn]bool

type data struct {
	Orders         OrderList
	Users          UserList
	Clients        map[string]websocketConnects
	BalanceChanges BalanceChanges
}

var Data = data{
	Orders:         OrderList{},
	Users:          UserList{},
	BalanceChanges: BalanceChanges{},
	Clients: map[string]websocketConnects{
		"orders": make(websocketConnects),
		"matchs": make(websocketConnects),
	},
}
