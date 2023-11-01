package common

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const (
	configPath = "../port.config"
	dsn        = "postgres://postgres:secret@localhost:%s?sslmode=disable"
)

func InitializeDatabase() *gorm.DB {

	port := mustGetPortFromPortConfig()

	cfg := fmt.Sprintf(dsn, port)

	db, err := gorm.Open(postgres.Open(cfg), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	db.AutoMigrate(&Product{})
	return db
}

func mustGetPortFromPortConfig() string {
	for {
		port, err := os.ReadFile(configPath)
		if err == nil {
			return string(port)
		}
		time.Sleep(time.Second)
	}
}
