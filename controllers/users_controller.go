package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/emj365/xschange/models"
)

func PostUsers(w http.ResponseWriter, r *http.Request) {
	usersFromReq := models.UserList{}
	err := json.NewDecoder(r.Body).Decode(&usersFromReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	for _, user := range usersFromReq {
		// add init balance after adding user
		balance := int(user.Balance)
		good := int(user.GoodAmount)
		user.Balance = 0
		user.GoodAmount = 0
		models.Data.Users = append(models.Data.Users, user)
		models.Data.BalanceChanges.Add(&models.BalanceChange{
			Good:              good,
			Balance:           balance,
			User:              user,
			IsDepositWithdrew: true,
		})
	}

	json.NewEncoder(w).Encode(models.Data.Users)
}
