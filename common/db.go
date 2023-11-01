package common

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	port = 49154
	dsn  = "postgres://postgres:secret@localhost:%d?sslmode=disable"
)

func InitializeDatabase() *gorm.DB {

	cfg := fmt.Sprintf(dsn, port)

	db, err := gorm.Open(postgres.Open(cfg), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	db.AutoMigrate(&Product{})
	return db
}
