package common

import (
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) String() string {
	return fmt.Sprintf("Product{ID: %d, Code: %s, Price: %d}", p.ID, p.Code, p.Price)
}

func CreateProduct(db *gorm.DB, code string, price uint) Product {
	p := Product{Code: code, Price: price}
	db.Create(&p)

	return p
}

func GetProduct(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProductByCode(db *gorm.DB, code string) (*Product, error) {
	var product Product
	if err := db.Where("code = ?", code).First(&product).Error; err != nil {
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
