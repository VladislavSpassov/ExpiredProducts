package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.Migrator().AutoMigrate(RecipyProduct{})
	db.Migrator().AutoMigrate(Product{})
	db.Migrator().AutoMigrate(Recipy{})
	DB = db
}
