package db

import (
	"github.com/mrtechit/purchase-transaction/model"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestStoreTrx(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Errorf("error connecting to DB")
	}
	defer teardown()

	newDB := NewDB(db)
	trx := &model.StoreTransaction{
		TransactionID:   "7804a376-3688-4187-a7ba-992893cd4cee",
		Description:     "test",
		TransactionDate: "10-11-2022",
		USDollarAmount:  "1.54",
	}
	err = newDB.StoreTrx(trx)
	require.NoError(t, err)
}

func setup() (*gorm.DB, error) {
	db, err := migrateAndConnectToPsgTestDB()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func teardown() {

}

func conntectToPsgTestDB() (*gorm.DB, error) {
	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrateAndConnectToPsgTestDB() (*gorm.DB, error) {

	db, err := conntectToPsgTestDB()
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
