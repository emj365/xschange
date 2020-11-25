package services

import (
	"encoding/json"
	"net/http"

	"github.com/emj365/xschange/models"
)

type Exchanger interface {
	// start()
	// stop()
	PlaceOrder(w http.ResponseWriter, o models.Order) error
	CancelOrder(w http.ResponseWriter, oID uint) error
}

type PublishEventFunc func(event models.Event)
type CreateMatch func(m models.Match)
type CreateBalanceChange func(m models.BalanceChange)

type OrdersManager interface {
	Place(o models.Order) error
	Cancel(oID uint) (models.Order, error)
	Process(PublishEventFunc, CreateMatch, CreateBalanceChange) error
}

type UsersManager interface {
	Find(uID uint) error
}

type MatchesManager interface {
	Create(m models.Match)
}

type BalanceChangesManager interface {
	Create(bc models.BalanceChange)
}

type Listener interface {
	Process(e models.Event)
}

/*
	ExchangeEngine
*/

var _ Exchanger = NewExchangeEngine(nil, nil, nil, nil, nil) // just for verification

func NewExchangeEngine(
	u UsersManager,
	o OrdersManager,
	m MatchesManager,
	bc BalanceChangesManager,
	l []Listener,
) *ExchangeEngine {
	return &ExchangeEngine{
		u, o, m, bc, l,
	}
}

type ExchangeEngine struct {
	Users          UsersManager
	Orders         OrdersManager
	Matches        MatchesManager
	BalanceChanges BalanceChangesManager
	listeners      []Listener
}

func (e ExchangeEngine) PlaceOrder(w http.ResponseWriter, o models.Order) error {
	if err := e.Users.Find(o.UserID); err != nil {
		return err
	}

	if err := e.Orders.Place(o); err != nil {
		return err
	}

	json.NewEncoder(w).Encode(o)

	return e.Orders.Process(
		e.publishEvent,
		e.Matches.Create,
		e.BalanceChanges.Create,
	)
}

func (e ExchangeEngine) CancelOrder(w http.ResponseWriter, oID uint) error {
	o, err := e.Orders.Cancel(oID)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(o)
	return nil
}

func (e ExchangeEngine) publishEvent(event models.Event) {
	for _, l := range e.listeners {
		l.Process(event)
	}
}
