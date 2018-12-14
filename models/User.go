package models

// User data
type User struct {
	Orders     OrderList `json:"-"`
	GoodAmount uint      `json:"goodAmount"`
	Balance    uint      `json:"balance"`
}

// UserAvaliableBalance indicate how much balance is avliable for the new order
type UserAvaliableBalance struct {
	GoodAmount uint
	Balance    uint
}

// GetAvaliableBalance calculate AvaliableBalance from the user's orders
func (u User) GetAvaliableBalance() UserAvaliableBalance {
	goodAmount := u.GoodAmount
	balance := u.Balance
	for _, o := range u.Orders {
		if o.Selling {
			goodAmount -= o.Remain
		} else {
			balance -= o.Remain * o.Price
		}
	}

	return UserAvaliableBalance{goodAmount, balance}
}

// CheckBalanceForOrder check user's avliable balance is enough for the order
func (u User) CheckBalanceForOrder(o Order) bool {
	avaliableBalance := u.GetAvaliableBalance()
	if o.Selling {
		return avaliableBalance.GoodAmount >= o.Quantity
	}

	return avaliableBalance.Balance >= o.Quantity*o.Price
}

// UserList contain users
type UserList []*User
