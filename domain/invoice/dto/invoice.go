package dto

import "github.com/asentientbanana/pausalkole-admin/models"

type InvoiceItemDto struct {
	Amount      float32 `json:"amount"`
	Quantity    int     `json:"quantity"`
	Metric      string  `json:"metric"`
	Description string  `json:"description"`
}

type InvoiceCurrency struct {
	Placement string `json:"placement" required:"true"`
	Symbol    string `json:"symbol" required:"true"`
	Value     string `json:"value" required:"true"`
	Label     string `json:"label" required:"true"`
}

type AddInvoiceDto struct {
	Agency        string           `json:"agency" binding:"required"`
	Recipient     string           `json:"recipient" binding:"required"`
	Description   string           `json:"description" binding:"required"`
	DateCompleted int              `json:"date_completed"`
	Items         []InvoiceItemDto `json:"items" binding:"required"`
	Total         int              `json:"total" binding:"required"`
	DateDue       int              `json:"date_due" binding:"required"`
	Currency      string           `json:"currency" binding:"required"`
	Status        string           `json:"status" binding:"required"`
}

type UpdateInvoiceDto struct {
	AddInvoiceDto
	ID string `json:"id" binding:"required"`
}

type CompleteInvoiceDto struct {
	models.Invoice
	Recipient models.Entity `json:"recipient" binding:"required"`
	Agency    models.Entity `json:"agency" binding:"required"`
}
