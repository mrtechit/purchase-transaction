package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mrtechit/purchase-transaction/db"
	"github.com/mrtechit/purchase-transaction/httpserver"
	"github.com/mrtechit/purchase-transaction/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {

	configDir := os.Getenv("CONFIG_DIR")

	if err := godotenv.Load(path.Join(configDir, ".env")); err != nil {
		fmt.Println(".env file will not be used")
	}
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
	dsn := os.Getenv("POSTGRES_DSN")
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
