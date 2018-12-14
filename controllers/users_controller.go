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

	models.Data.Users = append(models.Data.Users, usersFromReq...)
	json.NewEncoder(w).Encode(models.Data.Users)
}
