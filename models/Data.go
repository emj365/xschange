package models

type data struct {
	Orders OrderList
	Users  UserList
}

var Data = data{
	Orders: OrderList{},
	Users: UserList{
		&User{Balance: 100, GoodAmount: 100},
		&User{Balance: 100, GoodAmount: 100},
	},
}
