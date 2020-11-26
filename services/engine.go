package services

import (
	"errors"
)

type Marketer interface {
	PlaceOrder(o Order) error
	CancelOrder(oID uint) error
}

/*
TradeExchangeEngine is implementation of Marketer
*/

var _ Marketer = NewExchangeEngine(nil, nil, nil, nil, nil)

func NewExchangeEngine(
	u UserExistentChecker,
	o Trader,
	m Matcher,
	ex Exchanger,
	l []Listener,
) *TradeExchangeEngine {
	return &TradeExchangeEngine{
		u, o, m, ex, l,
	}
}

type TradeExchangeEngine struct {
	userExistentChecker UserExistentChecker
	trader              Trader
	matcher             Matcher
	exchanger           Exchanger
	listeners           []Listener
}

func (e TradeExchangeEngine) PlaceOrder(o Order) error {
	if e.userExistentChecker.Check(o.UserID) {
		return errors.New("user not found")
	}

	e.trader.Place(o)
	orders := e.trader.GetOrders()
	e.publishEvent(Event{Name: "order has placed", Payload: o})

	matches := e.matcher.Match(&o, orders)
	e.exchanger.Exchange(o.UserID, matches, func(c BalanceChange) {
		e.publishEvent(Event{Name: "balance has changed", Payload: c})
	})

	return nil
}

func (e TradeExchangeEngine) CancelOrder(oID uint) error {
	o, canceled := e.trader.Cancel(oID)
	if !canceled {
		return errors.New("order can not be cancel ")
	}

	e.publishEvent(Event{Name: "order has canceled", Payload: o})
	return nil
}

func (e TradeExchangeEngine) publishEvent(event Event) {
	for _, l := range e.listeners {
		l.Listen(event)
	}
}

/*
dependencies of TradeExchangeEngine
*/

type UserExistentChecker interface {
	Check(uID uint) bool
}

type Trader interface {
	Place(o Order)
	Cancel(oID uint) (Order, bool)
	GetOrders() *[]Order
}

type Matcher interface {
	Match(o *Order, orders *[]Order) *[]Match
}

type Exchanger interface {
	Exchange(uID uint, o *[]Match, onChanged func(BalanceChange))
}

type Listener interface {
	Listen(e Event)
}

/*
models
*/

type Event struct {
	Name    string
	Payload interface{}
}

type Order struct {
	UserID    uint     `json:"userID"`
	Selling   bool     `json:"selling"`
	Quantity  uint     `json:"quantity"`
	Remain    uint     `json:"remain"`
	Price     uint     `json:"price"`
	Matchs    []*Match `json:"-"`
	CreatedAt int64    `json:"createAt"`
}

// Match is one of trades of order
type Match struct {
	Order    *Order
	Quantity uint
	Price    uint
}

type User struct {
	Orders     []*Order `json:"-"`
	GoodAmount uint     `json:"goodAmount"`
	Balance    uint     `json:"balance"`
}

// UserAvaliableBalance indicate how much balance is avaliable for the new order
type UserAvaliableBalance struct {
	GoodAmount uint
	Balance    uint
}

type BalanceChange struct {
	Good              int
	Balance           int
	User              *User
	Match             *Match // be provided if it is trade
	Order             *Order // be provided if it is margin
	IsDepositWithdrew bool
}
