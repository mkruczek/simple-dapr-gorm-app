package main

import (
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, code string, price uint) {
	db.Create(&Product{Code: code, Price: price})
}

func GetProduct(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(db *gorm.DB, id uint, newCode string, newPrice uint) {
	product, err := GetProduct(db, id)
	if err != nil {
		return
	}
	product.Code = newCode
	product.Price = newPrice
	db.Save(&product)
}

func DeleteProduct(db *gorm.DB, id uint) {
	db.Delete(&Product{}, id)
}
