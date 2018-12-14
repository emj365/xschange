package models

// Match is one of trades of order
type Match struct {
	Order    *Order
	Quantity uint
	Price    uint
}

// MatchList is list of matched orders
type MatchList []*Match

// ExchangeAssets handle user assets changes by matched orders
func (matchs *MatchList) ExchangeAssets(orderUserID uint, users *[]*User) {
	for _, m := range *matchs {
		forBuyerOrder := m.Order.Selling

		orderUser := (*users)[int(orderUserID)]
		matchedUser := (*users)[int(m.Order.UserID)]

		var buyer, seller *User
		if forBuyerOrder {
			buyer, seller = orderUser, matchedUser
		} else {
			buyer, seller = matchedUser, orderUser
		}

		buyer.GoodAmount += m.Quantity
		seller.GoodAmount -= m.Quantity

		amount := m.Quantity * m.Price
		buyer.Balance -= amount
		seller.Balance += amount
	}
}
