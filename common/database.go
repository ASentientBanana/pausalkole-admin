package common

import (
	"fmt"
	"github.com/asentientbanana/pausalkole-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=pgdb sslmode=disable password=%s port=5433", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	fmt.Println(dsn)
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
