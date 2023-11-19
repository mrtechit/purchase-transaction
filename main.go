package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
