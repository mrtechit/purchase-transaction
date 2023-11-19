package model

type StoreTransaction struct {
	Description     string `gorm:"primaryKey"`
	TransactionDate string
	USDollarAmount  string
}
