package models

import (
	"sort"
	"time"
)

type OrderList []*Order

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
	peers = filterByType(&peers, !o.Selling)
	peers = filterByPrice(&peers, !o.Selling, o.Price)
	sortPeers(&peers, !o.Selling)

	o.LinkMatchedOrders(&peers)
	o.Matchs.ExchangeAssets(o.UserID, users)
	*orders = append(*orders, o)
}

// private

func filterByType(peers *OrderList, forBuyer bool) OrderList {
	var new OrderList
	for _, o := range *peers {
		unclosed := o.Remain != 0
		if o.Selling == forBuyer && unclosed {
			new = append(new, o)
		}
	}

	return new
}

func filterByPrice(peers *OrderList, forBuyer bool, price uint) OrderList {
	var new OrderList
	for _, p := range *peers {
		goodPriceForBuyer := p.Price <= price
		goodPriceForSeller := p.Price >= price

		if forBuyer && goodPriceForBuyer {
			new = append(new, p)
		}

		if !forBuyer && goodPriceForSeller {
			new = append(new, p)
		}
	}

	return new
}

func sortPeers(peers *OrderList, selling bool) {
	forBuyer := !selling
	sort.Slice(*peers, func(i, j int) bool {
		currentGreatThanNext := (*peers)[i].Price > (*peers)[j].Price

		if forBuyer {
			return currentGreatThanNext
		}

		return !currentGreatThanNext
	})
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
