package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

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

func GetAuditProduct(db *gorm.DB, id uint) (*AuditProduct, error) {
	var AuditProduct AuditProduct
	if err := db.First(&AuditProduct, id).Error; err != nil {
		return nil, err
	}
	return &AuditProduct, nil
}

func GetAuditProductByCode(db *gorm.DB, code string) (*AuditProduct, error) {
	var AuditProduct AuditProduct
	if err := db.Where("code = ?", code).First(&AuditProduct).Error; err != nil {
		return nil, err
	}
	return &AuditProduct, nil
}

func UpdateAuditProduct(db *gorm.DB, id uint, newCode string, newPrice uint) {
	AuditProduct, err := GetAuditProduct(db, id)
	if err != nil {
		return
	}
	AuditProduct.Code = newCode
	AuditProduct.Price = newPrice
	db.Save(&AuditProduct)
}

func DeleteAuditProduct(db *gorm.DB, id uint) {
	db.Delete(&AuditProduct{}, id)
}
