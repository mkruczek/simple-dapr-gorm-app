package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"simple-gorm-app/common"
	"time"
)

func main() {

	db := common.InitializeDatabase()
	mustInitProductTable(db)

	http.HandleFunc("/audit/products", createAuditProductHandler(db))

	log.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func mustInitProductTable(db *gorm.DB) {
	err := db.AutoMigrate(&AuditProduct{})
	if err != nil {
		log.Fatal("cannot migrate audit product table: ", err)
	}
}

func createAuditProductHandler(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
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

type AuditProduct struct {
	gorm.Model
	Code      string
	Price     uint
	StoreDate time.Time
}

func (ap *AuditProduct) String() string {
	return fmt.Sprintf("AuditProduct{ID: %d, Code: %s, Price: %d}", ap.ID, ap.Code, ap.Price)
}

func CreateAuditProduct(db *gorm.DB, code string, price uint) AuditProduct {
	p := AuditProduct{Code: code, Price: price, StoreDate: time.Now()}
	db.Create(&p)

	return p
}
