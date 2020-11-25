package services

import (
	"net/http"

	"github.com/emj365/xschange/models"
)

type Engine interface {
	// start()
	// stop()
	PlaceOrder(w http.ResponseWriter, o models.Order) error
	CancelOrder(oID uint) error
	// addEventListener(listener Listener)
}

/*
  repos
*/

// orders

type machingPublisher interface {
	PublishMatching(models.Match)
}

type OrdersRepo interface {
	Place(o models.Order) error
	Cancel(oID uint) error
	Process(machingPublisher)
}

// users

type UsersRepo interface {
	Find(uID uint) error
}

// balance changes

type BalanceChangesRepo interface {
}

/*
	libs
*/

type Publisher interface {
	PublishMatching(models.Match)
}

/*
	XsEngine
*/

var _ Engine = NewXsEngine(nil, nil, nil, nil)

func NewXsEngine(p Publisher, u UsersRepo, o OrdersRepo, bc BalanceChangesRepo) XsEngine {
	return XsEngine{
		p, u, o, bc,
	}
}

type XsEngine struct {
	Publisher      Publisher
	Users          UsersRepo
	Orders         OrdersRepo
	BalanceChanges BalanceChangesRepo
}

func (e XsEngine) PlaceOrder(w http.ResponseWriter, o models.Order) error {
	if err := e.Users.Find(o.UserID); err != nil {
		return err
	}

	if err := e.Orders.Place(o); err != nil {
		return err
	}

	go e.Orders.Process(e.Publisher)
	return nil
}

func (e XsEngine) CancelOrder(oID uint) error {
	return e.Orders.Cancel(oID)
}
