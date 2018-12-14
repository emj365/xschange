package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order", func() {

	Describe("Order.LinkMatchedOrders", func() {

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
