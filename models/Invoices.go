package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Amount      float32
	Quantity    int
	Metric      string
	Description string
	InvoiceID   string
}

type InvoiceCurrencies struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	Label     string
	Value     string
	Symbol    string
	Placement string
}

type Invoice struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primaryKey;type:uuid"`
	DateCompleted int
	Recipient     string `gorm:"foreignKey:UserID"`
	Agency        string `gorm:"foreignKey:UserID"`
	Total         int
	DateDue       int
	Description   string
	Currency      string `gorm:"foreignKey:InvoiceCurrencies"`
	Status        string
	Items         []InvoiceItem `gorm:"foreignKey:InvoiceID"`
	UserID        string
}
