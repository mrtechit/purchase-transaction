package db

import (
	"github.com/mrtechit/purchase-transaction/model"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestStoreTrx_Success(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Errorf("error connecting to DB")
	}
	trx := &model.StoreTransaction{
		TransactionID:   "7804a376-3688-4187-a7ba-992893cd4cee",
		Description:     "test",
		TransactionDate: "10-11-2022",
		USDollarAmount:  "1.54",
	}
	defer teardown(db, *trx)

	newDB := NewDB(db)

	err = newDB.StoreTrx(trx)
	require.NoError(t, err)
}

func TestStoreTrx_Failed_DuplicateKey(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Errorf("error connecting to DB")
	}
	trx := &model.StoreTransaction{
		TransactionID:   "7804a376-3688-4187-a7ba-992893cd4cmm",
		Description:     "test",
		TransactionDate: "10-11-2022",
		USDollarAmount:  "1.54",
	}
	defer teardown(db, *trx)

	newDB := NewDB(db)

	err = newDB.StoreTrx(trx)
	err = newDB.StoreTrx(trx)
	require.Error(t, err)
}

func TestRetrieveTrx_Success(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Errorf("error connecting to DB")
	}
	trxStruct := &model.StoreTransaction{
		TransactionID:   "7804a376-3688-4187-a7ba-992893cd4cdf",
		Description:     "test",
		TransactionDate: "10-11-2022",
		USDollarAmount:  "1.54",
	}
	defer teardown(db, *trxStruct)
	newDB := NewDB(db)
	err = newDB.StoreTrx(trxStruct)
	if err != nil {
		t.Errorf("duplicate record")
	}

	trx, err := newDB.RetrieveTrx(trxStruct.TransactionID)
	require.NoError(t, err)
	require.Equal(t, "1.54", trx.USDollarAmount)
}

func TestRetrieveTrx_Failed_TrxNotFound(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Errorf("error connecting to DB")
	}
	trxStruct := &model.StoreTransaction{
		TransactionID:   "7804a376-3688-4187-a7ba-992893cddddd",
		Description:     "test",
		TransactionDate: "10-11-2022",
		USDollarAmount:  "1.54",
	}
	defer teardown(db, *trxStruct)
	newDB := NewDB(db)

	_, err = newDB.RetrieveTrx(trxStruct.TransactionID)
	require.Error(t, err)
}

func setup() (*gorm.DB, error) {
	db, err := migrateAndConnectToPsgTestDB()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func teardown(db *gorm.DB, transaction model.StoreTransaction) {
	db.Delete(&transaction)
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
