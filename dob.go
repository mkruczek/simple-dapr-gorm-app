package main

import (
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func InitializeDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Nie można połączyć się z bazą danych")
	}
	db.AutoMigrate(&Product{})
	return db
}
