package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func DbConnect() *gorm.DB {
	dsn := "docker:pass@tcp(golang-db:3306)/main?charset=utf8mb4&parseTime=true&Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("OpenDB failed", err)
	}
	return db
}