package models

import (
	"time"
)

// Order created by the user
type Order struct {
	UserID    uint      `json:"userID"`
	Selling   bool      `json:"selling"`
	Quantity  uint      `json:"quantity"`
	Remain    uint      `json:"remain"`
	Price     uint      `json:"price"`
	Matchs    MatchList `json:"-"`
	CreatedAt int64     `json:"createAt"`
}

// Place create a new order
func (o *Order) Place(orders *OrderList, users *UserList) {
	o.Remain = o.Quantity
	o.CreatedAt = time.Now().UnixNano()

	peers := *orders
	peers.FilterByType(!o.Selling).FilterByPrice(!o.Selling, o.Price).Sort(!o.Selling)

	o.LinkMatchedOrders(&peers)
	o.Matchs.ExchangeAssets(o.UserID, users)
	*orders = append(*orders, o)
}

// LinkMatchedOrders set remain for both the order & matched orders and create Matchs for the order
func (o *Order) LinkMatchedOrders(matchedOrders *OrderList) {
	for _, matchedOrder := range *matchedOrders {
		var matchedQuantity uint
		var closeOrders, uncloseOrders OrderList

		peerRemainExactlyMatch := matchedOrder.Remain == o.Remain
		peerRemainIsGreater := matchedOrder.Remain > o.Remain

		if peerRemainExactlyMatch {
			closeOrders = append(closeOrders, o, matchedOrder)
		} else if peerRemainIsGreater {
			matchedQuantity = o.Remain
			closeOrders = append(closeOrders, o)
			uncloseOrders = append(uncloseOrders, matchedOrder)
		} else {
			matchedQuantity = matchedOrder.Remain
			closeOrders = append(closeOrders, matchedOrder)
			uncloseOrders = append(uncloseOrders, o)
		}

		// change reamin to 0 for orders those suppose to close
		for _, closeOrder := range closeOrders {
			closeOrder.Remain = 0
		}

		// calculate remain for orders those can be closed
		for _, uncloseOrder := range uncloseOrders {
			uncloseOrder.Remain -= matchedQuantity
		}

		// create Match with matched quantity of peer
		match := Match{Order: matchedOrder,
			Quantity: matchedQuantity, Price: matchedOrder.Price}
		o.Matchs = append(o.Matchs, &match)

		// break if this order can be closed
		if peerRemainExactlyMatch || peerRemainIsGreater {
			break
		}
	}
}
