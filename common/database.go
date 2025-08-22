package common

import (
	"github.com/asentientbanana/pausalkole-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=pgdb user=pausalkole password=ADMIN1234! port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.Users{}, &models.Invoice{}, &models.InvoiceItem{}, &models.Entity{}, &models.EntityField{}, &models.InvoiceCurrencies{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
