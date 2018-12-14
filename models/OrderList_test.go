package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderList", func() {

	Describe("OrderList.FilterByType", func() {
		Context("with orders", func() {
			orders := OrderList{
				&Order{Selling: true, Quantity: 1, Remain: 1},
				&Order{Selling: true, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 0},
			}

			It("should filter pair orders by type", func() {
				result := orders.FilterByType(false)
				Expect(len(*result)).To(Equal(3))
			})
		})
	})

	Describe("OrderList.Sort", func() {
		Context("with orders have price", func() {
			orders := OrderList{
				&Order{Price: 11},
				&Order{Price: 9},
				&Order{Price: 12},
				&Order{Price: 10},
			}

			It("should filter sort ASC for buyer", func() {
				orders.Sort(false)
				Expect(orders[0].Price).To(Equal(uint(12)))
				Expect(orders[1].Price).To(Equal(uint(11)))
				Expect(orders[2].Price).To(Equal(uint(10)))
				Expect(orders[3].Price).To(Equal(uint(9)))
			})

			It("should filter sort DESC for seller", func() {
				orders.Sort(true)
				Expect(orders[0].Price).To(Equal(uint(9)))
				Expect(orders[1].Price).To(Equal(uint(10)))
				Expect(orders[2].Price).To(Equal(uint(11)))
				Expect(orders[3].Price).To(Equal(uint(12)))
			})
		})
	})

	Describe("OrderList.FilterByPrice", func() {
		Context("with orders have price", func() {
			orders := OrderList{
				&Order{Price: 1},
				&Order{Price: 2},
				&Order{Price: 3},
				&Order{Price: 4},
				&Order{Price: 5},
				&Order{Price: 6},
			}

			It("filter for orders with pair price for buyer", func() {
				result := *orders.FilterByPrice(true, 4)
				Expect(len(result)).To(Equal(4))
				Expect(result[0].Price).To(Equal(uint(1)))
				Expect(result[1].Price).To(Equal(uint(2)))
				Expect(result[2].Price).To(Equal(uint(3)))
				Expect(result[3].Price).To(Equal(uint(4)))
			})

			It("filter for orders with pair price for seller", func() {
				result := *orders.FilterByPrice(false, 4)
				Expect(len(result)).To(Equal(3))
				Expect(result[0].Price).To(Equal(uint(4)))
				Expect(result[1].Price).To(Equal(uint(5)))
				Expect(result[2].Price).To(Equal(uint(6)))
			})
		})
	})

})
