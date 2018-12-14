package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order", func() {
	Describe("filterByType", func() {
		Context("with orders", func() {

			peers := OrderList{
				&Order{Selling: true, Quantity: 1, Remain: 1},
				&Order{Selling: true, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 0},
			}

			It("should filter pair orders by type", func() {
				result := filterByType(&peers, false)
				Expect(len(result)).To(Equal(3))
			})
		})
	})

	Describe("sortPeers", func() {
		Context("with orders have price", func() {

			peers := OrderList{
				&Order{Price: 11},
				&Order{Price: 9},
				&Order{Price: 12},
				&Order{Price: 10},
			}

			It("should filter sort ASC for buyer", func() {
				sortPeers(&peers, false)
				Expect(peers[0].Price).To(Equal(uint(12)))
				Expect(peers[1].Price).To(Equal(uint(11)))
				Expect(peers[2].Price).To(Equal(uint(10)))
				Expect(peers[3].Price).To(Equal(uint(9)))
			})

			It("should filter sort DESC for seller", func() {
				sortPeers(&peers, true)
				Expect(peers[0].Price).To(Equal(uint(9)))
				Expect(peers[1].Price).To(Equal(uint(10)))
				Expect(peers[2].Price).To(Equal(uint(11)))
				Expect(peers[3].Price).To(Equal(uint(12)))
			})
		})
	})

	Describe("filterByPrice", func() {
		Context("with orders have price", func() {

			peers := OrderList{
				&Order{Price: 1},
				&Order{Price: 2},
				&Order{Price: 3},
				&Order{Price: 4},
				&Order{Price: 5},
				&Order{Price: 6},
			}

			It("filter for orders with pair price for buyer", func() {
				result := filterByPrice(&peers, true, 4)
				Expect(len(result)).To(Equal(4))
				Expect(result[0].Price).To(Equal(uint(1)))
				Expect(result[1].Price).To(Equal(uint(2)))
				Expect(result[2].Price).To(Equal(uint(3)))
				Expect(result[3].Price).To(Equal(uint(4)))
			})

			It("filter for orders with pair price for seller", func() {
				result := filterByPrice(&peers, false, 4)
				Expect(len(result)).To(Equal(3))
				Expect(result[0].Price).To(Equal(uint(4)))
				Expect(result[1].Price).To(Equal(uint(5)))
				Expect(result[2].Price).To(Equal(uint(6)))
			})
		})
	})

	Describe("match", func() {
		Context("with peers for buyer", func() {

			peers := OrderList{
				&Order{Remain: 2, Price: 2},
				&Order{Remain: 1, Price: 2},
				&Order{Remain: 3, Price: 3},
				&Order{Remain: 1, Price: 4},
			}

			It("create match for buyer", func() {
				order := Order{Remain: 4, Price: 3, Selling: false}
				order.LinkMatchedOrders(&peers)

				Expect(len(order.Matchs)).To(Equal(3))

				Expect(order.Matchs[0].Order).To(Equal(peers[0]))
				Expect(order.Matchs[0].Quantity).To(Equal(uint(2)))
				Expect(order.Matchs[0].Price).To(Equal(uint(2)))

				Expect(order.Matchs[1].Order).To(Equal(peers[1]))
				Expect(order.Matchs[1].Quantity).To(Equal(uint(1)))
				Expect(order.Matchs[1].Price).To(Equal(uint(2)))

				Expect(order.Matchs[2].Order).To(Equal(peers[2]))
				Expect(order.Matchs[2].Quantity).To(Equal(uint(1)))
				Expect(order.Matchs[2].Price).To(Equal(uint(3)))

				Expect(order.Remain).To(Equal(uint(0)))
				Expect(peers[0].Remain).To(Equal(uint(0)))
				Expect(peers[1].Remain).To(Equal(uint(0)))
				Expect(peers[2].Remain).To(Equal(uint(2)))
			})
		})

		Context("with peers for seller", func() {

			peers := OrderList{
				&Order{Remain: 1, Price: 4},
				&Order{Remain: 3, Price: 3},
				&Order{Remain: 2, Price: 2},
				&Order{Remain: 1, Price: 2},
			}

			It("create match for buyer", func() {
				order := Order{Remain: 5, Price: 4, Selling: true}
				order.LinkMatchedOrders(&peers)

				Expect(len(order.Matchs)).To(Equal(3))

				Expect(order.Matchs[0].Order).To(Equal(peers[0]))
				Expect(order.Matchs[0].Quantity).To(Equal(uint(1)))
				Expect(order.Matchs[0].Price).To(Equal(uint(4)))

				Expect(order.Matchs[1].Order).To(Equal(peers[1]))
				Expect(order.Matchs[1].Quantity).To(Equal(uint(3)))
				Expect(order.Matchs[1].Price).To(Equal(uint(3)))

				Expect(order.Matchs[2].Order).To(Equal(peers[2]))
				Expect(order.Matchs[2].Quantity).To(Equal(uint(1)))
				Expect(order.Matchs[2].Price).To(Equal(uint(2)))

				Expect(order.Remain).To(Equal(uint(0)))
				Expect(peers[0].Remain).To(Equal(uint(0)))
				Expect(peers[1].Remain).To(Equal(uint(0)))
				Expect(peers[2].Remain).To(Equal(uint(1)))
			})
		})
	})
})
