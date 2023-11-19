package main

import (
	"github.com/mrtechit/purchase-transaction/db"
	"github.com/mrtechit/purchase-transaction/httpserver"
	"github.com/mrtechit/purchase-transaction/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	gormDb, err := migrateAndConnectToPsg()
	if err != nil {
		log.Fatal("[main] error init database")
	}
	dbManager := db.NewDB(gormDb)

	handler := httpserver.NewApiHandler(dbManager)
	handler.Handler()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("[main] error init http server")
	}
}

func conntectToPsg() (*gorm.DB, error) {
	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrateAndConnectToPsg() (*gorm.DB, error) {

	db, err := conntectToPsg()
	if err != nil {
		return nil, err
	}

	// Perform database migration
	err = db.AutoMigrate(&model.StoreTransaction{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
