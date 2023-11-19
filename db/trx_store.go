package db

import (
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
