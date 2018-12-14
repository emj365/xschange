package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	Describe("User.GetAvaliableBalance", func() {
		Context("with orders", func() {
			orders := OrderList{
				&Order{Selling: true, Quantity: 2, Remain: 2, Price: 10},
				&Order{Selling: true, Quantity: 2, Remain: 0, Price: 10},
				&Order{Selling: false, Quantity: 2, Remain: 2, Price: 10},
				&Order{Selling: false, Quantity: 2, Remain: 0, Price: 10},
			}

			user := User{Orders: orders, GoodAmount: uint(100), Balance: uint(100)}

			It("should filter pair orders by type", func() {
				UserAvaliableBalance := user.GetAvaliableBalance()
				Expect(UserAvaliableBalance.GoodAmount).To(Equal(uint(98)))
				Expect(UserAvaliableBalance.Balance).To(Equal(uint(80)))
			})
		})
	})

	Describe("CheckBalanceForOrder", func() {
		Context("with orders", func() {
			orders := OrderList{
				&Order{Selling: true, Quantity: 2, Remain: 2, Price: 10},
				&Order{Selling: true, Quantity: 2, Remain: 0, Price: 10},
				&Order{Selling: false, Quantity: 2, Remain: 2, Price: 10},
				&Order{Selling: false, Quantity: 2, Remain: 0, Price: 10},
			}

			user := User{Orders: orders, GoodAmount: uint(100), Balance: uint(100)}

			It("check balance for selling", func() {
				Expect(user.CheckBalanceForOrder(Order{Selling: false, Quantity: 2, Price: 41})).To(Equal(false))
				Expect(user.CheckBalanceForOrder(Order{Selling: false, Quantity: 2, Price: 40})).To(Equal(true))
			})

			It("check balance for buying", func() {
				Expect(user.CheckBalanceForOrder(Order{Selling: true, Quantity: 99, Price: 999})).To(Equal(false))
				Expect(user.CheckBalanceForOrder(Order{Selling: true, Quantity: 98, Price: 999})).To(Equal(true))
			})
		})
	})

})
