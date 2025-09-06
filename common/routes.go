package common

import (
	"github.com/asentientbanana/pausalkole-admin/domain/auth"
	"github.com/asentientbanana/pausalkole-admin/domain/currency"
	"github.com/asentientbanana/pausalkole-admin/domain/entity"
	"github.com/asentientbanana/pausalkole-admin/domain/invoice"
	"github.com/asentientbanana/pausalkole-admin/domain/pdf"
	"github.com/asentientbanana/pausalkole-admin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(server *gin.Engine, db *gorm.DB) {

	protected := server.Group("/").Use(middleware.AuthMiddleware())

	// Auth
	server.POST("/auth/login", func(context *gin.Context) {
		auth.HandleLogin(context, db)
	})

	// User
	server.POST("/user/register", func(context *gin.Context) {
		auth.HandleRegister(context, db)
	})

	// Entity
	protected.POST("/entities", func(context *gin.Context) {
		entity.AddEntity(context, db)
	})
	protected.PUT("/entities", func(context *gin.Context) {
		entity.UpdateEntity(context, db)
	})
	protected.GET("/entities", func(context *gin.Context) {
		entity.GetEntities(context, db)
	})
	protected.DELETE("/entities/:id", func(context *gin.Context) {
		id := context.Param("id")
		entity.DeleteEntity(context, db, id)
	})

	// Invoice
	protected.GET("/invoices", func(context *gin.Context) {
		invoice.GetAllUserInvoices(context, db)
	})
	protected.POST("/invoices", func(context *gin.Context) {
		invoice.AddInvoice(context, db)
	})
	protected.DELETE("/invoices/:id", func(context *gin.Context) {
		invoice.DeleteInvoice(context, db, context.Param("id"))
	})
	protected.PUT("/invoices/:id", func(context *gin.Context) {
		invoice.UpdateInvoice(context, db)
	})
	protected.GET("/invoices/currencies", func(context *gin.Context) {
		currency.GetCurrencies(context, db)
	})
	protected.GET("/invoices/document/:id", func(context *gin.Context) {
		pdf.GetInvoicePdf(context, db, context.Param("id"))
	})
}
