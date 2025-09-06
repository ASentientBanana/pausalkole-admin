package models

import (
	"github.com/google/uuid"
)

type InvoiceItem struct {
	BaseModel
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Amount      float32   `gorm:"default:0" json:"amount"`
	Quantity    int       `json:"quantity"`
	Metric      string    `json:"metric"`
	Description string    `json:"description"`
	InvoiceID   string    `json:"invoice_id"`
}

type InvoiceCurrencies struct {
	BaseModel
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	Symbol    string    `json:"symbol"`
	Placement string    `json:"placement"`
}

type Invoice struct {
	BaseModel
	ID            uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	DateCompleted int
	// Foreign keys
	RecipientID uuid.UUID `gorm:"type:uuid" json:"recipient_id"`
	AgencyID    uuid.UUID `gorm:"type:uuid" json:"agency_id"`

	// Both reference Entity table
	Recipient Entity `gorm:"foreignKey:RecipientID;references:ID" json:"recipient"`
	Agency    Entity `gorm:"foreignKey:AgencyID;references:ID" json:"agency"`

	Total       int           `gorm:"default:0" json:"total"`
	DateDue     int           `gorm:"default:0" json:"date_due"`
	Description string        `gorm:"type:text" json:"description"`
	Currency    string        `gorm:"foreignKey:InvoiceCurrencies" json:"currency"`
	Status      string        `gorm:"type:text" json:"status"`
	Items       []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
	UserID      string        `gorm:"type:uuid" json:"user_id"`
}
