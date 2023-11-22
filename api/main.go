package main

import (
	"log"
	"net/http"
	"simple-gorm-app/common"
)

func main() {

	db := common.InitializeDatabase()

	http.HandleFunc("/products", CreateProduct(db))

	log.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
