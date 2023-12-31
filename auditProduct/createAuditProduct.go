package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func CreateAuditProductHandler(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
			return
		}

		var payload AuditProduct
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := CreateAuditProduct(db, payload.Code, payload.Price)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}
}
