package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MatchList", func() {
	Describe("ExchangeAssets", func() {
		Context("for seller order", func() {
			users := UserList{
				&User{Balance: 10, GoodAmount: 10},
				&User{Balance: 10, GoodAmount: 10},
				&User{Balance: 10, GoodAmount: 10},
				&User{Balance: 10, GoodAmount: 10},
			}

			matchs := MatchList{
				&Match{Order: &Order{Selling: false, UserID: 1}, Quantity: 1, Price: 2},
				&Match{Order: &Order{Selling: false, UserID: 2}, Quantity: 2, Price: 1},
				&Match{Order: &Order{Selling: false, UserID: 3}, Quantity: 3, Price: 1},
			}

			order := Order{UserID: 0, Matchs: matchs}

			It("exchange assets", func() {
				order.Matchs.ExchangeAssets(order.UserID, &users)

				Expect(users[0].Balance).To(Equal(uint(17)))
				Expect(users[1].Balance).To(Equal(uint(8)))
				Expect(users[2].Balance).To(Equal(uint(8)))
				Expect(users[3].Balance).To(Equal(uint(7)))
				Expect(users[0].GoodAmount).To(Equal(uint(4)))
				Expect(users[1].GoodAmount).To(Equal(uint(11)))
				Expect(users[2].GoodAmount).To(Equal(uint(12)))
				Expect(users[3].GoodAmount).To(Equal(uint(13)))
			})
		})

		Context("for buyer order", func() {
			users := UserList{
				&User{Balance: 10, GoodAmount: 10},
				&User{Balance: 10, GoodAmount: 10},
				&User{Balance: 10, GoodAmount: 10},
				&User{Balance: 10, GoodAmount: 10},
			}

			matchs := MatchList{
				&Match{Order: &Order{Selling: true, UserID: 1}, Quantity: 1, Price: 2},
				&Match{Order: &Order{Selling: true, UserID: 2}, Quantity: 2, Price: 1},
				&Match{Order: &Order{Selling: true, UserID: 3}, Quantity: 3, Price: 1},
			}

			order := Order{UserID: 0, Matchs: matchs}

			It("exchange assets", func() {
				order.Matchs.ExchangeAssets(order.UserID, &users)

				Expect(users[0].GoodAmount).To(Equal(uint(16)))
				Expect(users[1].GoodAmount).To(Equal(uint(9)))
				Expect(users[2].GoodAmount).To(Equal(uint(8)))
				Expect(users[3].GoodAmount).To(Equal(uint(7)))
				Expect(users[0].Balance).To(Equal(uint(3)))
				Expect(users[1].Balance).To(Equal(uint(12)))
				Expect(users[2].Balance).To(Equal(uint(12)))
				Expect(users[3].Balance).To(Equal(uint(13)))
			})
		})
	})
})
