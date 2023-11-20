package db

import (
	"errors"
	"github.com/mrtechit/purchase-transaction/model"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func NewDB(Db *gorm.DB) *DB {
	return &DB{Db: Db}
}

func (db *DB) StoreTrx(trx *model.StoreTransaction) error {
	result := db.Db.Create(&trx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DB) RetrieveTrx(transactionID string) (*model.StoreTransaction, error) {
	var storeTransaction model.StoreTransaction
	storeTransaction.TransactionID = transactionID
	result := db.Db.Find(&storeTransaction)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("trx not found")
	}
	return &storeTransaction, nil
}
