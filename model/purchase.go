package model

type StoreTransaction struct {
	TransactionID   string `gorm:"primaryKey"`
	Description     string
	TransactionDate string
	USDollarAmount  string
}
