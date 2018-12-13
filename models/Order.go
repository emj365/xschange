package models

import (
	"sort"
	"time"
)

// Order created by the user
type Order struct {
	UserID    uint     `json:"userID"`
	Selling   bool     `json:"selling"`
	Quantity  uint     `json:"quantity"`
	Remain    uint     `json:"remain"`
	Price     uint     `json:"price"`
	Matchs    []*Match `json:"-"`
	CreatedAt int64    `json:"createAt"`
}

// Place create a new order
func (o *Order) Place(orders *[]*Order, users *[]*User) {
	o.Remain = o.Quantity
	o.CreatedAt = time.Now().UnixNano()

	peers := *orders
	peers = filterByType(&peers, !o.Selling)
	peers = filterByPrice(&peers, !o.Selling, o.Price)
	sortPeers(&peers, !o.Selling)

	o.match(orders, &peers)
	o.doSettlement(users)
	*orders = append(*orders, o)
}

// private

func filterByType(peers *[]*Order, forBuyer bool) []*Order {
	var new []*Order
	for _, o := range *peers {
		unclosed := o.Remain != 0
		if o.Selling == forBuyer && unclosed {
			new = append(new, o)
		}
	}

	return new
}

func filterByPrice(peers *[]*Order, forBuyer bool, price uint) []*Order {
	var new []*Order
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

func sortPeers(peers *[]*Order, selling bool) {
	forBuyer := !selling
	sort.Slice(*peers, func(i, j int) bool {
		currentGreatThanNext := (*peers)[i].Price > (*peers)[j].Price

		if forBuyer {
			return currentGreatThanNext
		}

		return !currentGreatThanNext
	})
}

// Match create matchs for the order
func (o *Order) match(orders, peers *[]*Order) {
	for _, p := range *peers {
		var matchedQuantity uint
		var closeOrders, uncloseOrders []*Order

		peerRemainExactlyMatch := p.Remain == o.Remain
		peerRemainIsGreater := p.Remain > o.Remain

		if peerRemainExactlyMatch {
			closeOrders = append(closeOrders, o, p)
		} else if peerRemainIsGreater {
			matchedQuantity = o.Remain
			closeOrders = append(closeOrders, o)
			uncloseOrders = append(uncloseOrders, p)
		} else {
			matchedQuantity = p.Remain
			closeOrders = append(closeOrders, p)
			uncloseOrders = append(uncloseOrders, o)
		}

		for _, closeOrder := range closeOrders {
			closeOrder.Remain = 0
		}

		for _, uncloseOrder := range uncloseOrders {
			uncloseOrder.Remain -= matchedQuantity
		}

		match := Match{Order: p, Quantity: matchedQuantity, Price: p.Price}
		o.Matchs = append(o.Matchs, &match)

		if peerRemainExactlyMatch || peerRemainIsGreater {
			break
		}
	}
}

// DoSettlement caculate and set users' balance
func (o *Order) doSettlement(users *[]*User) {
	var buyer, seller *User
	if o.Selling {
		seller = (*users)[int(o.UserID)]
	} else {
		buyer = (*users)[int(o.UserID)]
	}

	for _, m := range (*o).Matchs {
		if o.Selling {
			buyer = (*users)[int((*m).Order.UserID)]
		} else {
			seller = (*users)[int((*m).Order.UserID)]
		}

		amount := m.Quantity * m.Price
		(*buyer).Balance -= amount
		(*seller).Balance += amount
	}
}
