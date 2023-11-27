package main

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"simple-gorm-app/common"
)

func main() {

	db := common.InitializeDatabase()
	initProductTable(db)

	http.HandleFunc("/audit/products", CreateAuditProductHandler(db))

	log.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func initProductTable(db *gorm.DB) {
	db.AutoMigrate(&AuditProduct{})
}
