package models

import (
	"sort"
)

type OrderList []*Order

func (orderList *OrderList) FilterByType(forBuyer bool) *OrderList {
	var new OrderList
	for _, o := range *orderList {
		unclosed := o.Remain != 0
		if o.Selling == forBuyer && unclosed {
			new = append(new, o)
		}
	}

	return &new
}

func (orderList *OrderList) FilterByPrice(forBuyer bool, price uint) *OrderList {
	var new OrderList
	for _, order := range *orderList {
		goodPriceForBuyer := order.Price <= price
		goodPriceForSeller := order.Price >= price

		if forBuyer && goodPriceForBuyer {
			new = append(new, order)
		}

		if !forBuyer && goodPriceForSeller {
			new = append(new, order)
		}
	}

	return &new
}

func (orderList *OrderList) Sort(selling bool) {
	forBuyer := !selling
	sort.Slice(*orderList, func(i, j int) bool {
		currentGreatThanNext := (*orderList)[i].Price > (*orderList)[j].Price

		if forBuyer {
			return currentGreatThanNext
		}

		return !currentGreatThanNext
	})
}
