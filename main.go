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

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Error loading env")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
