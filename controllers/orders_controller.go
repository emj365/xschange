package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emj365/xschange/libs"
	"github.com/emj365/xschange/models"
)

func extractOrderFromRequest(r *http.Request, o *models.Order) error {
	err := json.NewDecoder(r.Body).Decode(o)
	return err
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Data.Orders)
}

func PostOrders(w http.ResponseWriter, r *http.Request) {
	o := models.Order{}
	err := extractOrderFromRequest(r, &o)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if err = o.Place(); err != nil {
		switch err {
		case models.UserNotExistErr:
			w.WriteHeader(http.StatusBadRequest)
			break
		case models.BalanceNotEnoughErr:
			w.WriteHeader(http.StatusNotAcceptable)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(o)

	libs.LogOrders(&models.Data.Orders)
	for i, u := range models.Data.Users {
		log.Printf("users[%v]: %v\n", i, *u)
	}
}
