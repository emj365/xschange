package models

import (
	"fmt"
	"log"
	"sort"
	"time"
)

// Order created by the user
type Order struct {
	UserID    uint     `json:"userID"`
	Selling   bool     `json:"selling"`
	Quantity  int      `json:"quantity"`
	Remain    int      `json:"remain"`
	Price     int      `json:"price"`
	Matchs    []*Match `json:"-"`
	CreatedAt int64    `json:"createAt"`
}

// DoSettlement caculate and set users' balance
func (o *Order) DoSettlement() {
}

// Match create matchs for the order
func (o *Order) Match(orders *[]*Order) {
	peers := *orders
	peers = filterByType(&peers, !o.Selling)
	peers = filterByPrice(&peers, !o.Selling, o.Price)
	sortPeers(&peers, !o.Selling)

	for _, peer := range peers {
		// fmt.Println(peer)
		var match Match
		if o.Remain >= peer.Remain {
			match = Match{Order: peer, Quantity: peer.Remain, Price: peer.Price}
			o.Remain -= peer.Remain
			peer.Remain = 0
		} else {
			match = Match{Order: peer, Quantity: o.Remain, Price: peer.Price}
			peer.Remain -= o.Remain
			o.Remain = 0
		}

		o.Matchs = append(o.Matchs, &match)
		if o.Remain == 0 {
			break
		}
	}
}

// Place create a new order
func (o *Order) Place(orders *[]*Order) {
	o.Remain = o.Quantity
	o.CreatedAt = time.Now().Unix()
	o.Match(orders)
	o.DoSettlement()
	*orders = append(*orders, o)

	fmt.Println("\033[2J")
	log.Printf("orders: %#v\n\n", orders)
	for i, o := range *orders {
		log.Printf("orders[%v]: %#v\n", i, *o)
		for j, p := range (*o).Matchs {
			log.Printf("orders[%v].Matchs[%v]: %#v\n", i, j, *p)
		}
	}
}

// private

func filterByType(peers *[]*Order, forBuyer bool) []*Order {
	var new []*Order
	for _, o := range *peers {
		if o.Selling == forBuyer && o.Remain != 0 {
			new = append(new, o)
		}
	}

	return new
}

func filterByPrice(peers *[]*Order, forBuyer bool, price int) []*Order {
	var new []*Order
	for _, p := range *peers {
		if forBuyer && p.Price <= price {
			new = append(new, p)
		}

		if !forBuyer && p.Price >= price {
			new = append(new, p)
		}
	}

	return new
}

func sortPeers(peers *[]*Order, selling bool) {
	sort.Slice(*peers, func(i, j int) bool {
		if selling {
			return (*peers)[i].Price < (*peers)[j].Price
		}

		return (*peers)[i].Price > (*peers)[j].Price
	})
}
