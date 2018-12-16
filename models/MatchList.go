package models

// MatchList is list of matched orders
type MatchList []*Match

// ExchangeAssets handle user assets changes by matched orders
func (matchs *MatchList) ExchangeAssets(orderUserID uint) {
	for _, m := range *matchs {
		forBuyerOrder := m.Order.Selling

		orderUser := Data.Users[int(orderUserID)]
		matchedUser := Data.Users[int(m.Order.UserID)]

		var buyer, seller *User
		if forBuyerOrder {
			buyer, seller = orderUser, matchedUser
		} else {
			buyer, seller = matchedUser, orderUser
		}

		amount := m.Quantity * m.Price

		// balance changes
		buyerBalanceChange := new(BalanceChange)
		buyerBalanceChange.Match = m
		buyerBalanceChange.User = buyer
		buyerBalanceChange.Good = int(m.Quantity)
		buyerBalanceChange.Balance = int(amount) * -1

		sellerBalanceChange := new(BalanceChange)
		sellerBalanceChange.Match = m
		sellerBalanceChange.User = seller
		sellerBalanceChange.Good = int(m.Quantity) * -1
		sellerBalanceChange.Balance = int(amount)

		Data.BalanceChanges.Add(buyerBalanceChange)
		Data.BalanceChanges.Add(sellerBalanceChange)
	}
}
