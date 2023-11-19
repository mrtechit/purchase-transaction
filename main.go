package main

import (
	"github.com/mrtechit/purchase-transaction/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrateDB() *gorm.DB {

	db, err := connectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	// Perform database migration
	err = db.AutoMigrate(&model.StoreTransaction{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
