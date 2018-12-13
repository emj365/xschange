package models

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Xschange Models")
}

var _ = Describe("Order", func() {
	Describe("choosePeers", func() {
		Context("with orders", func() {

			peers := []*Order{
				&Order{Selling: true, Quantity: 1, Remain: 1},
				&Order{Selling: true, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 1},
				&Order{Selling: false, Quantity: 1, Remain: 0},
			}

			It("should filter pair orders by type", func() {
				result := choosePeers(&peers, false)
				Expect(len(result)).To(Equal(3))
			})
		})
	})

	Describe("sortPeers", func() {
		Context("with orders have price", func() {

			peers := []*Order{
				&Order{Price: 11},
				&Order{Price: 9},
				&Order{Price: 12},
				&Order{Price: 10},
			}

			It("should filter sort ASC for buyer", func() {
				sortPeers(&peers, false)
				Expect(peers[0].Price).To(Equal(12))
				Expect(peers[1].Price).To(Equal(11))
				Expect(peers[2].Price).To(Equal(10))
				Expect(peers[3].Price).To(Equal(9))
			})

			It("should filter sort DESC for seller", func() {
				sortPeers(&peers, true)
				Expect(peers[0].Price).To(Equal(9))
				Expect(peers[1].Price).To(Equal(10))
				Expect(peers[2].Price).To(Equal(11))
				Expect(peers[3].Price).To(Equal(12))
			})
		})
	})

	Describe("filterByPrice", func() {
		Context("with orders have price", func() {

			peers := []*Order{
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
				Expect(result[0].Price).To(Equal(1))
				Expect(result[1].Price).To(Equal(2))
				Expect(result[2].Price).To(Equal(3))
				Expect(result[3].Price).To(Equal(4))
			})

			It("filter for orders with pair price for seller", func() {
				result := filterByPrice(&peers, false, 4)
				Expect(len(result)).To(Equal(3))
				Expect(result[0].Price).To(Equal(4))
				Expect(result[1].Price).To(Equal(5))
				Expect(result[2].Price).To(Equal(6))
			})
		})
	})
})
