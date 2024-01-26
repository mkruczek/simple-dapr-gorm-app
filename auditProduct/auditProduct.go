package main

import (
	"gorm.io/gorm"
)

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
